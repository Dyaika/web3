package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

func main() {
	loadEnv()
	http.HandleFunc("/process", processHandler)

	port := getEnv("PORT", "8080")

	log.Printf("Server is running on port %s...\n", port)
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
	// Получаем данные из Cookie
	if encoded, err := r.Cookie(cookieName); err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode(cookieName, encoded.Value, &value); err == nil {
			newResponse = "Previous Request: " + value["Request"]
		}
	}

	// Обработка нового запроса
	var data UserData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Сохранение данных в Cookie
	if encoded, err := cookieHandler.Encode(cookieName, map[string]string{
		"Request":  data.Request,
		"Response": data.Response,
	}); err == nil {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    cookieName,
			Value:   encoded,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
	}

	// Отправка ответа
	w.Write([]byte("Request Processed\n" + newResponse))
}
