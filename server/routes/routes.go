package routes

import (
	"landingpage/auth"
	"landingpage/campaign"
	"landingpage/config"
	"landingpage/middleware"
	"landingpage/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(r *gin.Engine) {
	db = config.GetDB()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := user.NewHandler(userService, authService)

	authMiddleware := middleware.AuthMiddleware(userService, authService)

	v1 := r.Group("/api/v1")
	router := v1

	userRoutes := router.Group("/users")
	userRoutes.POST("/", userHandler.RegisterUser)
	userRoutes.POST("/sessions", userHandler.LoginUser)
	userRoutes.POST("/email_checkers", userHandler.CheckEmailAvailability)
	userRoutes.POST("/upload_avatar", authMiddleware, userHandler.UploadAvatar)

	// router.Use(authMiddleware)
	CampaignRoutes(router)
}

func CampaignRoutes(router *gin.RouterGroup) {
	campaignRoutes := router.Group("/campaigns")

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := campaign.NewHandler(campaignService)

	campaignRoutes.GET("/", campaignHandler.FindCampaigns)
}
