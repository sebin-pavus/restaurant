package web

import (
	"restaurant/internal/web/handler"

	"github.com/gin-gonic/gin"
)

func NewServer(router *gin.Engine) {
	router.GET("/top_items", handler.GetTopMenuItems)
}
