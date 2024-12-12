package controllers

import (
	"fmt"
	"net/http"

	"github.com/chris-weir/chrisweir/views"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	view, err := views.ParseFromFile(views.FS, "404.gohtml", "app.gohtml")
	if err != nil {
		fmt.Println("Failed to parse from file", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	view.Execute(w, r, nil)
}
