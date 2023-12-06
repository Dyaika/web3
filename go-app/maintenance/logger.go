package maintenance

import (
	"log"
	"os"
	"time"
)

func LogData(message string) {
	log.Println(message)
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("AYOO")
	}
	logFile.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
	logFile.Close()
}
