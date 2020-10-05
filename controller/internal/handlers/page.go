package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct{
	HttpClient *http.Client
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
	return
}

func (h *Handler) Web(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/form.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

