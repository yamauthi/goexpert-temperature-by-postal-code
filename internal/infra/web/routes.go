package web

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}
