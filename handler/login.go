package handler

import (
	"log"
	"net/http"
	"text/template"
)

type LoginFormError struct {
	UserName string
	Password string
}
type LoginStudent struct {
	UserName  string
	Password  string
	FormError LoginFormError
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	parseLoginTemplate(w, LoginStudent{})
}

func (h Handler) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%v", err)
	}
	un := r.PostFormValue("UserName")
	pass := r.PostFormValue("Password")

	if un == "" {
		parseLoginTemplate(w, LoginStudent{
			UserName: un,
			Password: "",
			FormError: LoginFormError{
				UserName: "The username field is required",
			},
		})
		return
	}
	if pass == "" {
		parseLoginTemplate(w, LoginStudent{
			UserName: un,
			FormError: LoginFormError{
				Password: "The password field is required",
			},
		})
		return
	}
	if un != "admin" {
		parseLoginTemplate(w, LoginStudent{
			UserName: un,
			FormError: LoginFormError{
				UserName: "The username/password doesn't match",
			},
		})
		return
	}
	if pass != "12345" {
		parseLoginTemplate(w, LoginStudent{
			UserName: un,
			FormError: LoginFormError{
				UserName: "The username/password doesn't match",
			},
		})
		return
	}
	h.sessionManager.Put(r.Context(), "username", un)
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

func parseLoginTemplate(w http.ResponseWriter, data any) {
	t, _ := template.ParseFiles("templates/login.html")
	if err := t.Execute(w, data); err != nil {
		log.Fatalf("%v", err)
	}
}
