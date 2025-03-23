package main

func registerRoutes(userController *UserController) {
	// Example:
	// http.HandleFunc("/register", userController.RegisterUser)
}

func main() {
	userModel := NewUserModel()
	userController := NewUserController(userModel)

	registerRoutes(userController)
}
