package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

func main() {
	http.HandleFunc("/process", processHandler)
	http.ListenAndServe(":8080", nil)
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
	// Получаем данные из Cookie
	var newResponse string = ""
	if encoded, err := r.Cookie("userdata"); err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("userdata", encoded.Value, &value); err == nil {
			// Ваши действия с данными, например, отправка ответа
			newResponse = "Previous Request: " + value["Request"]
		}
	}

	// Обработка нового запроса
	var data UserData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ваша логика обработки данных
	data.Response = "Your Response"

	// Сохранение данных в Cookie
	if encoded, err := cookieHandler.Encode("userdata", map[string]string{
		"Request":  data.Request,
		"Response": data.Response,
	}); err == nil {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    "userdata",
			Value:   encoded,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
	}

	// Отправка ответа
	w.Write([]byte("Request Processed\n" + newResponse))
}
