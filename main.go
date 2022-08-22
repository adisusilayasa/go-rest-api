package main

import (
	articleRepo "belajar-interface/article/repository"
	articleService "belajar-interface/article/service"
	"belajar-interface/database"
	"belajar-interface/route"
	userRepo "belajar-interface/user/repository"
	userService "belajar-interface/user/service"
	"log"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Print(err.Error())
	}
	articleRepository := articleRepo.NewRepository(db)
	articleService := articleService.NewService(articleRepository)
	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)

	route.NewHandler(articleService, userService)

}
