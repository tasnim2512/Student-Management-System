package handler

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
)

func (h Handler) EditStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roll, _ := strconv.Atoi(id)
	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	var editStudent Student
	for _, val := range ul.Students {
		if val.ID == roll {
			editStudent = val
			break
		}
	}

	parseEditTemplate(w, editStudent)
}
func (h Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Sid,  _ := strconv.Atoi(id)

	student := Student{}
	err := h.decoder.Decode(&student, r.PostForm)
	if err != nil {
		log.Fatal(err)
	}

	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			student.FormError = vErr
		}
		parseEditTemplate(w, student)
		return
	}
	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}
	for idx, st := range ul.Students {
		if st.ID == Sid {
			ul.Students[idx].Name = student.Name
			ul.Students[idx].Email = student.Email
			ul.Students[idx].Roll = student.Roll
			break
		}
	}
	if err := writeUsersToFile(ul); err != nil {
		log.Fatalf("%v", err)
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

func parseEditTemplate(w http.ResponseWriter, data any) {
	t, _ := template.ParseFiles("templates/edit-Student.html")
	if err := t.Execute(w, data); err != nil {
		log.Fatalf("%v", err)
	}
}
