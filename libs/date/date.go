package date

import (
	"time"
)

func GetCurrentFormattedDate() string {
	currentDateTime := time.Now()
	return currentDateTime.Format("01/02/2006 15:04:05")
}
