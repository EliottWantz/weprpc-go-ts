package main

import (
	"log"
	"net/http"

	"test/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	webrpcHandler := service.NewUserServiceServer(service.NewUserService())
	r.Handle("/*", webrpcHandler)

	log.Fatal(http.ListenAndServe(":8081", r))
}
