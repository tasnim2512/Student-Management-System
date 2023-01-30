package handler

import (
	"fmt"
	"net/http"
)

func  (h Handler) HOME(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
