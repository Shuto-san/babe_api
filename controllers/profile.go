package controllers

func ShowProfiles (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a profile handler!")
}

func ShowUserProfiles (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a profile handler!")
}

func storeProfiles (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "this is a profile handler!")
}
