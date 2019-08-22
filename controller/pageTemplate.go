package controller

import (
	"EFServer/tool"
	"html/template"
)

var indexTemplate = template.Must(template.New("index").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/index.html"))))

type indexVM struct {
	IsLogin  bool
	UserName string
}

var loginInputTemplate = template.Must(template.New("login").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/loginInput.html"))))

type loginVM struct {
	Tip string
}

var userRegistInputTemplate = template.Must(template.New("regist").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/registInput.html"))))

type registInputVM struct {
	Tip string
}

var loginSuccessTemplate = template.Must(template.New("login").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/loginSuccess.html"))))
var userRegistSuccessTemplate = template.Must(template.New("regist").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/registSuccess.html"))))
