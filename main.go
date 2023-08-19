package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	campaignService := campaign.NewService(campaignRepository)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	campaigns, _ := campaignService.FindCampaigns(2)
	fmt.Println(len(campaigns))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	router.Run()
}

/*
1. ambil nilai header authorization: bearer tokentokentoken
2. dari header authorization, kita ambil nilai tokennya saja
3. kita validasi token
4. ambil user_id
5. ambil user daru db berdasarkan user_id lewat service
6. set context isinya user
*/
