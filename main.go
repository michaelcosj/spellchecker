package main

import (
	"embed"
	_ "embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// wordlist obtained from https://www.mit.edu/~ecprice/wordlist.10000
// wordlist.2 obtained from https://github.com/dwyl/english-words
//
//go:embed assets/wordlist.2
var wordlist string

//go:embed template/*
var content embed.FS

//go:embed assets/public/*
var public embed.FS

func main() {
	dictionary := strings.Fields(string(wordlist))
	service := &Service{dictionary}

	t, err := template.ParseFS(content, "template/*.tmpl.html")
	HandleError(err)

	handler := &Handler{service, t}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.Index)
	r.Post("/spellcheck", handler.SpellCheck)

	publicFs, err := fs.Sub(public, "assets/public")
	HandleError(err)

	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(publicFs))))

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
