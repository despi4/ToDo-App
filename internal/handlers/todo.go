package handlers

import (
	"fmt"
	"net/http"
)

func Todo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is todo page")
}
