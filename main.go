package main

import (
    "github.com/Shuto-san/babe_api/routes"
    "net/http"
)

func main() {

    apiRouter := routes.SetupApiRouter()

    http.ListenAndServe(":80", apiRouter)
}
