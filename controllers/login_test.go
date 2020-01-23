package controllers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/Shuto-san/babe_api/controllers"
)

func TestLogin(t *testing.T) {

    var body = []byte(`{"id"}:"1234567890"`)

    req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.Login)

    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"returnUrl" : "/babe"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
