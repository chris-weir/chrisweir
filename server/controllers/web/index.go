package controllers

import (
	"fmt"
	"net/http"

	"github.com/chris-weir/chrisweir/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view, err := views.ParseFromFile(views.FS, "index.gohtml", "app.gohtml")
	if err != nil {
		fmt.Println("Failed to parse from file", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	view.Execute(w, r, nil)
}
