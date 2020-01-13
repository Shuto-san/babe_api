package routes

import (
  "github.com/gorilla/mux"
)

func SetupApiRouter() {

  apiRouter := mux.NewRouter()

  apiRouter.HandleFunc("/register", controllers.Register).Methods("POST")
  apiRouter.HandleFunc("/login", controllers.Login).Methods("POST")
  apiRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")
  apiRouter.HandleFunc("/questions", controllers.ShowQuestions).Methods("GET")
  apiRouter.HandleFunc("/profiles", controllers.ShowProfiles).Methods("GET")
  apiRouter.HandleFunc("/profiles", controllers.StoreProfiles).Methods("POST")
  apiRouter.HandleFunc("/users/{userId}/profiles", controllers.ShowUserProfiles).Methods("GET")
  apiRouter.HandleFunc("/users/recommended", controllers.ShowRecommendedUsers).Methods("GET")
  apiRouter.HandleFunc("/users/{userId}/impression", controllers.StoreUserImpression).Methods("POST")
  apiRouter.HandleFunc("/reviews", controllers.ShowReviews).Methods("GET")
  apiRouter.HandleFunc("/users/{userId}/reviews", controllers.ShowUserReviews).Methods("GET")
  apiRouter.HandleFunc("/users/{userId}/review", controllers.StoreUserReview).Methods("POST")
  apiRouter.HandleFunc("/reviews/{reviewId}/block", controllers.BlockReview).Methods("POST")
  apiRouter.HandleFunc("/chatRoomConfigs", controllers.ShowChatRoomConfigs).Methods("GET")
  apiRouter.HandleFunc("/chatRoomConfigs", controllers.StoreChatRoomConfigs).Methods("POST")
  apiRouter.HandleFunc("/users/{userId}/chatRoomConfigs/block", controllers.BlockChatRoomConfig).Methods("POST")

  return apiRouter
}
