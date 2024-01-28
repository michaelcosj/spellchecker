package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// wordlist obtained from https://www.mit.edu/~ecprice/wordlist.10000
	data, err := os.ReadFile("./assets/wordlist")
	HandleError(err)

	dictionary := strings.Fields(string(data))
	service := &Service{dictionary}

	t, err := template.ParseGlob("./templates/*.html")
	HandleError(err)
	handler := &Handler{service, t}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.Index)
	r.Post("/spellcheck", handler.SpellCheck)

	fs := http.FileServer(http.Dir("./assets/public/"))
	r.Handle("/*", http.StripPrefix("/", fs))

	log.Printf("Starting server...\n")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	http.ListenAndServe(":"+port, r)
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
