package handler

import (
	"to-do-list/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{

		profile := api.Group("/profile")
		{
			profile.GET("/", h.getUserProfile)
			profile.PUT("/", h.updateUserProfile)
		}

		chat := api.Group("/chat")
		{
			chat.POST("/send-message", h.sendMessage)
			chat.GET("/messages", h.getMessages)
			chat.POST("/create-chat-room", h.createChatRoom)
			chat.POST("/add-user-to-room", h.addUserToChatRoom)
			chat.POST("/remove-user-from-room", h.removeUserFromChatRoom)
		}
	}
	return router
}
