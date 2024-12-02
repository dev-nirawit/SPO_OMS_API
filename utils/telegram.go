// utils/telegram.go
package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// TelegramBot โครงสร้างสำหรับ Telegram Bot
type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

// NewTelegramBot ฟังก์ชันสำหรับสร้าง instance ของ TelegramBot
func NewTelegramBot() (*TelegramBot, error) {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	// ดึงค่า TELEGRAM_APITOKEN จาก .env
	botToken := os.Getenv("TELEGRAM_APITOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_APITOKEN is not set in the environment")
	}

	// เริ่มต้นใช้งาน bot ด้วย token ที่ดึงมาได้
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("Error creating Telegram bot: %v", err)
	}

	// เปิด debug mode เพื่อดู log
	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	// ดึง Chat ID จาก environment variable
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing TELEGRAM_CHAT_ID: %v", err)
	}

	return &TelegramBot{
		Bot:    bot,
		ChatID: chatID,
	}, nil
}

// SendMessage ส่งข้อความไปยัง Telegram
func (t *TelegramBot) SendMessage(chatID int64, msgText string) {
	msg := tgbotapi.NewMessage(chatID, msgText)
	if _, err := t.Bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

// HandleCommand ฟังก์ชันสำหรับจัดการคำสั่ง
func (t *TelegramBot) HandleCommand(command string, chatID int64) {
	switch command {
	case "ChatID":
		t.SendMessage(chatID, fmt.Sprintf("Your Chat ID is: %d", chatID))
	default:
		t.SendMessage(chatID, "คำสั่งไม่รู้จัก กรุณาลองใหม่อีกครั้ง")
	}
}

// HandleMessage ฟังก์ชันสำหรับจัดการข้อความที่ไม่ใช่คำสั่ง
func (t *TelegramBot) HandleMessage(message string, chatID int64) {
	// ตอบกลับข้อความที่ถูกสอบถาม
	t.SendMessage(chatID, fmt.Sprintf("คุณถามว่า: %s", message))
}

// ListenForUpdates ฟังก์ชันเพื่อฟังการอัปเดตจาก Telegram
func (t *TelegramBot) ListenForUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ถ้ามีข้อความใหม่
			chatID := update.Message.Chat.ID
			messageText := update.Message.Text
			// fmt.Sprintf("messageText: %s", messageText)

			if update.Message.IsCommand() {
				command := strings.TrimPrefix(messageText, "/") // เอา '/' ออก
				t.HandleCommand(command, chatID)
			} else if strings.HasPrefix(messageText, "สอบถาม") {
				// ตัดข้อความเพื่อกรองคำถาม
				trimmedMessage := strings.TrimSpace(strings.TrimPrefix(messageText, "สอบถาม"))
				if trimmedMessage != "" {
					t.HandleMessage(trimmedMessage, chatID) // ส่งข้อความที่กรองแล้วไปยัง HandleMessage
				} else {
					t.SendMessage(chatID, "กรุณาถามคำถามหลังจาก 'สอบถาม'")
				}
			} else {
				// การจัดการข้อความอื่น ๆ สามารถเพิ่มได้ที่นี่
				t.SendMessage(chatID, "คำสั่งไม่รู้จัก กรุณาลองใหม่อีกครั้ง")
			}
		}
	}
}
