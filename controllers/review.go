package controllers

func ShowReviews(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a review handler!")
}

func ShowUserReviews(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a review handler!")
}

func StoreUserReview(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a review handler!")
}

func BlockReview(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a review handler!")
}
