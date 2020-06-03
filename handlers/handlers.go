package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckCookie(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		c, err := r.Cookie("mycookie")
		if err != nil || c.Value != "123456" {
			log.Println("Err cookie is : ", err)
			json.NewEncoder(w).Encode(&RESP{Notice: "no cookie"})
			return
		}
		// everything is ok
		h.ServeHTTP(w, r)
	})
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	var req REQ
	json.NewDecoder(r.Body).Decode(&req)
	// log.Println(req.Username, " : ", req.Password)
	if req.Username == "manh" && req.Password == "1" {
		http.SetCookie(w, &http.Cookie{
			Name:  "mycookie",
			Value: "123456",
		})
		// log.Println("ok ban oi")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(&RESP{Notice: "manhdeptrai"})
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&RESP{Notice: "wrong"})
	}
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&RESP{Notice: "nothing"})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("i'm in welcome")
	w.Write([]byte("hello, u passed"))
}

type REQ struct {
	Username string `json:"name"`
	Password string `json:"pass"`
}

type RESP struct {
	Notice string `json:"notice"`
}
