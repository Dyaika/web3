package maintenance

import (
	"log"
	"os"
	"time"
)

// логирование
var logFile *os.File

// InitializeLogging Инициализация логирования
func InitializeLogging() {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file for writing")
	}

	// Устанавливаем вывод в файл
	log.SetOutput(logFile)
}

func LogData(message string) {
	log.Println(message)
	logFile.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
