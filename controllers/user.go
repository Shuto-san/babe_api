package controllers

import (
    "fmt"
    "net/http"
)

func ShowRecommendedUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a user handler!")
}

func StoreUserImpression(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a user handler!")
}
