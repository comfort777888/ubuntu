package handlers

import (
	"net/http"
)

// checkMethodAndPath func receives request , path and method to compare. As if request does not match with m
// method or path then function returns status code error. Otherwise returns success  - code 200.
func checkMethodAndPath(r *http.Request, path, method string) int {
	if r.Method != method {
		return http.StatusMethodNotAllowed // 405
	}
	if r.URL.Path != path {
		return http.StatusNotFound // 404
	}
	return 200
}
