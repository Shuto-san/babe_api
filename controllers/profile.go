package controllers

import (
    "fmt"
    "net/http"
)

func ShowProfiles (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a profile handler!")
}

func ShowUserProfiles (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a profile handler!")
}

func StoreProfiles (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a profile handler!")
}
