package main

import (
	"RestApi/router"
	"RestApi/users"
)

func main() {
	users.InitializeMigration()
	router.InitializeRouter()
}
