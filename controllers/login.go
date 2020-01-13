package controllers

func Login (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a login handler!")
}
