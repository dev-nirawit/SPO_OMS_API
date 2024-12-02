package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Enabled  bool
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}

type PostgresConfig struct {
	Enabled  bool
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}

var MySQLConnections map[string]*gorm.DB
var PostgresConnections map[string]*gorm.DB

// LoadEnv โหลด environment variables จากไฟล์ .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// ConnectDB เชื่อมต่อกับฐานข้อมูลทั้งหมดที่กำหนดใน .env
func ConnectDB() {
	LoadEnv()

	MySQLConnections = make(map[string]*gorm.DB)
	PostgresConnections = make(map[string]*gorm.DB)

	connectMySQL()
	connectPostgres()
}

// connectMySQL เชื่อมต่อกับ MySQL hosts
func connectMySQL() {
	for i := 0; ; i++ {
		enabledStr := os.Getenv(fmt.Sprintf("MYSQL_ENABLED[%d]", i))
		if enabledStr == "" {
			break
		}

		enabled, err := strconv.ParseBool(enabledStr)
		if err != nil || !enabled {
			continue
		}

		host := os.Getenv(fmt.Sprintf("MYSQL_HOSTS[%d]", i))
		user := os.Getenv(fmt.Sprintf("MYSQL_USER[%d]", i))
		password := os.Getenv(fmt.Sprintf("MYSQL_PASSWORD[%d]", i))
		db := os.Getenv(fmt.Sprintf("MYSQL_DB[%d]", i))
		port := os.Getenv(fmt.Sprintf("MYSQL_PORT[%d]", i))

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, db)

		dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to connect to MySQL host %s: %v", host, err)
			continue
		}

		MySQLConnections[host] = dbConnection
		fmt.Printf("MySQL connection established for host: %s\n", host)
	}
}

// connectPostgres เชื่อมต่อกับ PostgreSQL hosts
func connectPostgres() {
	for i := 0; ; i++ {
		enabledStr := os.Getenv(fmt.Sprintf("POSTGRES_ENABLED[%d]", i))
		if enabledStr == "" {
			break
		}

		enabled, err := strconv.ParseBool(enabledStr)
		if err != nil || !enabled {
			continue
		}

		host := os.Getenv(fmt.Sprintf("POSTGRES_HOSTS[%d]", i))
		user := os.Getenv(fmt.Sprintf("POSTGRES_USER[%d]", i))
		password := os.Getenv(fmt.Sprintf("POSTGRES_PASSWORD[%d]", i))
		db := os.Getenv(fmt.Sprintf("POSTGRES_DB[%d]", i))
		port := os.Getenv(fmt.Sprintf("POSTGRES_PORT[%d]", i))

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, db, port)

		dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to connect to PostgreSQL host %s: %v", host, err)
			continue
		}

		PostgresConnections[host] = dbConnection
		fmt.Printf("PostgreSQL connection established for host: %s\n", host)
	}
}
