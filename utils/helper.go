// utils/helpers.go
package utils

import (
	"fmt"
	"time"
)

// FormatThaiDateTime แปลงเวลาให้เป็นรูปแบบวันเวลาในไทย
func FormatThaiDateTime(t time.Time) string {
	// แปลงปีจาก ค.ศ. เป็น พ.ศ.
	thaiYear := t.Year() + 543
	monthNames := []string{
		"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
		"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
	}
	month := monthNames[t.Month()-1] // เดือนเริ่มต้นจาก 1

	// จัดรูปแบบวันที่เป็น "2 พฤศจิกายน 2567 15:30:00"
	return fmt.Sprintf("%d %s %d %02d:%02d:%02d",
		t.Day(), month, thaiYear, t.Hour(), t.Minute(), t.Second())
}
