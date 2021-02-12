package main

import (
	"strings"

	"golang.org/x/net/html"
)

var MyFuncMap = map[string]interface{}{
	"ToLower":      strings.ToLower,
	"ToUpper":      strings.ToUpper,
	"ToGetterName": ToGetterName,
	"ToSetterName": ToSetterName,
	"ToSelector":   ToSelector,
	"ToClassName":  ToClassName,
	"escapeHtml":   EscapeHtml,
	"escapeQuote":  EscapeQuote,
	"ToImport":     ToImport,
}

func ToGetterName(name string) string {
	return "get" + strings.Title(name)
}
func ToSetterName(name string) string {
	return "Set" + strings.Title(name)
}
func ToSelector(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	var first = string(name[2])
	if strings.ToLower(string(name[3])) == string(name[3]) {
		first = strings.ToLower(string(name[2]))
	}

	return first + name[3:] + "Validator"
}
func ToClassName(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	return strings.ToUpper(string(name[2])) + name[3:] + "ValidatorDirective"
}
func ToImport(name string) string {

	return name[:len(name)-3]
}

func EscapeHtml(name string) string {
	return html.EscapeString(name)
}
func EscapeQuote(name string) string {
	return strings.Replace(name, "'", "\\'", 1)
}
