package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type key int

const userIPKey key = 0

type User struct {
	Name  string
	Email string
}

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

func routes() {
	http.HandleFunc("/user", findUser)
}

func findUser(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	ch := make(chan *User, 1)
	go func() {
		if ip, ok := ctx.Value(userIPKey).(string); ok {
			log.Println(" Старт БД", ip)
		}

		ch <- readDatabase()
		log.Println("GoRouting с БД уничтожен")
	}()

	select {
	case u := <-ch:
		sendResponse(rw, u, http.StatusOK)
		log.Println("Отправлен статус ОК")
		return

	case <-ctx.Done():
		e := struct {
			Error string
		}{
			ctx.Err().Error(),
		}
		sendResponse(rw, e, http.StatusRequestTimeout)
		log.Println("Отправлен статус превышения ожидания")
		return
	}
}

func readDatabase() *User {
	u := User{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	time.Sleep(100 * time.Millisecond)

	return &u
}

func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}
