package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizda")

	select {
	case <-time.After(5 * time.Second):
		log.Println("requeste processada com sucesso")

		w.Write([]byte("Request procesada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}

}
