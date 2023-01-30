package handler

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type StudentList struct {
	Students []Student `json:"students"`
}
type Student struct {
	ID        int              `json:"id" form:"-"`
	Name      string           `json:"Name"`
	Roll      int64            `json:"Roll"`
	Email     string           `json:"Email"`
	Hobbies   []string         `json:"Hobbies"`
	FormError map[string]error `json:"-"`
}

func (s Student) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name,
			validation.Required.Error("the name field is required"),
			validation.Length(3, 32).Error("the name length should be between 3 to 32"),
			validation.By(CheckAlreadyExists),
		),
		validation.Field(&s.Email,
			validation.Required.Error("the email field is required"),
			is.Email.Error("Please put a Valid email"),
		),
		validation.Field(&s.Roll,
			validation.Required.Error("the roll field is required"),
		),
	)
}
func (h Handler) STUDENT(w http.ResponseWriter, r *http.Request) {
	ul, err := getUserList()
	if err != nil {
		log.Fatalf("%v", err)
	}

	t, _ := template.ParseFiles("templates/student-list.html")
	if err := t.Execute(w, ul); err != nil {
		log.Fatalf("%v", err)
	}
}
func getUserList() (*StudentList, error) {
	studentFile, err := os.Open("students.json")
	if err != nil {
		return nil, err
	}

	defer studentFile.Close()

	ul := StudentList{}
	content, err := io.ReadAll(studentFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(content, &ul); err != nil {
		return nil, err
	}

	return &ul, err
}

func writeUsersToFile(ul *StudentList) error {
	file, err := json.MarshalIndent(ul, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("students.json", file, 0644)
	if err != nil {
		return err
	}
	return err
}

func CheckAlreadyExists(value any) error {
	name, ok := value.(string)
	if !ok {
		return errors.New("The name should be string")
	}
	ul, err := getUserList()
	if err != nil {
		log.Print(err)
		return err
	}
	for _, student := range ul.Students {
		if student.Name == name {
			return errors.New("the name is already exits")
		}
	}
	return nil
}
