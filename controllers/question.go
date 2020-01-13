package controllers

func ShowQuestions (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a question handler!")
}
