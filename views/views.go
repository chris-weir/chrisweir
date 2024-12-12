package views

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"text/template"
)

type View struct {
	Template *template.Template
}

//go:embed *
var FS embed.FS

func (v View) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tpl, err := v.Template.Clone()
	if err != nil {
		http.Error(w, "Failed to render page.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to build page.", http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buffer)
}

func ParseFromFile(fs fs.FS, patterns ...string) (View, error) {
	htmlTemplate := template.New(filepath.Base(patterns[0]))

	htmlTemplate, err := htmlTemplate.ParseFS(fs, patterns...)
	if err != nil {
		return View{}, fmt.Errorf("failed to parse template: %s", err)
	}

	return View{Template: htmlTemplate}, nil
}
