package controllers

import (
    "fmt"
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a register handler!")
}
