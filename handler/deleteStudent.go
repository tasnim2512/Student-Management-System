package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h Handler) DeleteStudent (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	Sid, _ := strconv.Atoi(id)

	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}

	var student StudentList
	for _, val := range ul.Students {
		if val.ID == Sid {
			continue
		} else {
			student.Students = append(student.Students, val)
		}
	}

	if err := writeUsersToFile(&student); err != nil {
		log.Fatalf("%v", err)
	}

	http.Redirect(w, r, "/students", http.StatusSeeOther)
}