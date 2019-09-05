package controller

import (
	"EFServer/tool"
	"html/template"
	"net/http"
)

var errorTemplate = template.Must(template.New("error").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/error.html"))))

func sendErrorPage(w http.ResponseWriter, err string) {
	errorTemplate.ExecuteTemplate(w, "error", err)
}
