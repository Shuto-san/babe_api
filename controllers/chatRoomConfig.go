package controllers

import (
    "fmt"
    "net/http"
)

func ShowChatRoomConfigs (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a chat room config handler!")
}

func StoreChatRoomConfigs (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a chat room config handler!")
}

func BlockChatRoomConfig (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "this is a chat room config handler!")
}
