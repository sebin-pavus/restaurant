package web

import (
	"restaurant/internal/web/handler"
	"restaurant/internal/web/service"

	"github.com/gin-gonic/gin"
)

func NewServer(router *gin.Engine) {
	newHandler := handler.HandlerStruct{Service: service.ServiceStruct{}}
	router.GET("/top_items", newHandler.GetTopMenuItems)
}
