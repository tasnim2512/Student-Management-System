package handler

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
)

func  (h Handler) DetailStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	Sid, _ := strconv.Atoi(id)

	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}

	var editStudent Student
	for _, val := range ul.Students {
		if val.ID == Sid {
			editStudent = val
		}
	}

	t, _ := template.ParseFiles("templates/student-detail.html")
	if err := t.Execute(w, editStudent); err != nil {
		log.Fatalf("%v", err)
	}
}
