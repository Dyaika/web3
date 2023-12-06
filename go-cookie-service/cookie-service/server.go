package cookie_service

import (
	"encoding/json"
	"github.com/gorilla/securecookie"
	"go-cookie-service/maintenance"
	"net/http"
	"strconv"
	"time"
)

func Start() {
	port := maintenance.GetEnv("PORT", "8080")
	maintenance.LogData("Server is running on port " + port + "\n")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
		return
	}
	http.HandleFunc("/process", processHandler)
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

type UserData struct {
	Request  string
	Response string
}

func megahard() {
	time.Sleep(1000)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	var cookieName = maintenance.GetEnv("COOKIE_NAME", "userdata")
	var newResponse string = ""
	// Получаем

	switch r.Method {
	case http.MethodPost:
		// Обработка нового запроса
		megahard()
		var data UserData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			maintenance.LogData("Error decoding JSON: " + err.Error())
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
			maintenance.LogData("Error encoding cookie-service:" + err.Error())
		}
	case http.MethodGet:
		go megahard()
		if encoded, err := r.Cookie(cookieName); err == nil {
			value := make(map[string]string)
			if err = cookieHandler.Decode(cookieName, encoded.Value, &value); err == nil {
				newResponse = "Data: " + value["Request"] + " : " + value["Response"]
			} else {
				newResponse = "data corrupted"
				maintenance.LogData("Error decoding cookie-service:" + err.Error())
			}
		} else {
			newResponse = "no data"
		}

	}
	// Отправка ответа
	w.Write([]byte("Request Processed\n" + newResponse))
}
