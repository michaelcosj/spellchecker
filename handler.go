package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service   *Service
	templates *template.Template
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h *Handler) SpellCheck(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	if word == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	countQuery := r.FormValue("count")
	count, err := strconv.Atoi(countQuery)
	if err != nil || count <= 0 {
		http.Error(w, http.StatusText(400), 400)
	}

	suggestions := h.service.GetSuggestions(strings.TrimSpace(word), count)
    if err := h.templates.ExecuteTemplate(w, "result_list", suggestions); err != nil {
		http.Error(w, http.StatusText(500), 500)
    }
}
