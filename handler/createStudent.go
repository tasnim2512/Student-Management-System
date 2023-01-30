package handler

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (h Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	parseCreateTemplate(w, Student{})
}
func (h Handler) StoreStudent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%v", err)
	}
	student := Student{}
	err := h.decoder.Decode(&student, r.PostForm)
	if err != nil {
		log.Fatal(err)
	}

	roll, _ := strconv.ParseInt(r.FormValue("Roll")[0:], 10, 64);
	student.Roll = roll
	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			student.FormError = vErr
			log.Println(vErr)
		}
		parseCreateTemplate(w, student)
		return
	}

	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}

	lastStudent := Student{}
	if len(ul.Students) >= 1 {
		lastStudent = ul.Students[len(ul.Students)-1]
	}
	student.ID = lastStudent.ID + 1

	ul.Students = append(ul.Students, student)
	if err := writeUsersToFile(ul); err != nil {
		log.Fatalf("%v", err)
	}

	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

func parseCreateTemplate(w http.ResponseWriter, data any) {
	t, _ := template.ParseFiles("templates/create-student.html")
	if err := t.Execute(w, data); err != nil {
		log.Fatalf("%v", err)
	}
}
