package http

import (
	"Notifications/internal/notification"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc notification.UseCase) {
	h := NewHandler(uc)

	router.POST("/", h.CreateNotification)
	router.GET("/", h.GetAllNotifications)
	router.GET("/:id", h.GetNotificationsById)
	router.GET("/user/:id", h.GetNotificationsByUser)
	router.PUT("/:id", h.UpdateNotification)
	router.DELETE("/:id", h.DeleteNotification)
}
