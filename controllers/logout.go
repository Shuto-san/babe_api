package controllers

import (
    "fmt"
    "net/http"
)

func Logout (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a login handler!")
}
