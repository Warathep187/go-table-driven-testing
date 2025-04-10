package main

import (
	"practical-example/controllers"
	"practical-example/models"
)

func registerRoutes(userController *controllers.UserController) {
	// Example:
	// http.HandleFunc("/register", userController.RegisterUser)
}

func main() {
	userModel := models.NewUserModel()
	badUsernameModel := models.NewBadUsernameModel()
	userController := controllers.NewUserController(userModel, badUsernameModel)

	registerRoutes(userController)
}
