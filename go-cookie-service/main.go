package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
)

func main() {
	loadEnv()
	initializeLogging()
	http.HandleFunc("/process", processHandler)

	port := getEnv("PORT", "8080")
	logData("Server is running on port " + port + "\n")
	http.ListenAndServe(":"+port, nil)
}

// Получение .env
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Получение переменных из окружения
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

type UserData struct {
	Request  string
	Response string
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	var cookieName = getEnv("COOKIE_NAME", "userdata")
	var newResponse string = ""
	// Получаем

	switch r.Method {
	case http.MethodPost:
		// Обработка нового запроса
		var data UserData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logData("Error decoding JSON:" + err.Error())
			return
		}

		// Сохранение
		if encoded, err := cookieHandler.Encode(cookieName, map[string]string{
			"Request":  data.Request,
			"Response": "length=" + strconv.Itoa(len(data.Request)),
		}); err == nil {
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{
				Name:    cookieName,
				Value:   encoded,
				Expires: expiration,
			}
			http.SetCookie(w, &cookie)
		} else {
			logData("Error encoding cookie:" + err.Error())
		}
	case http.MethodGet:
		time.Sleep(time.Second * 2)
		if encoded, err := r.Cookie(cookieName); err == nil {
			value := make(map[string]string)
			if err = cookieHandler.Decode(cookieName, encoded.Value, &value); err == nil {
				newResponse = "Data: " + value["Request"] + " : " + value["Response"]
			} else {
				newResponse = "data corrupted"
				logData("Error decoding cookie:" + err.Error())
			}
		} else {
			newResponse = "no data"
		}

	}
	// Отправка ответа
	w.Write([]byte("Request Processed\n" + newResponse))
}

// логирование
var logFile *os.File

// Инициализация логирования
func initializeLogging() {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file for writing")
	}

	// Устанавливаем вывод в файл
	log.SetOutput(logFile)
}

func logData(message string) {
	log.Println(message)
	logFile.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
