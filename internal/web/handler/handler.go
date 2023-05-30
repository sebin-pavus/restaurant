package handler

import (
	"net/http"
	"os"

	"restaurant/internal/model"
	"restaurant/internal/web/service"

	"github.com/gin-gonic/gin"
)

func GetTopMenuItems(c *gin.Context) {
	logFile := os.Getenv("LOG_FILE")
	topMenuItems, err := service.GetTopMenuItems(logFile)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	c.IndentedJSON(http.StatusOK, topMenuItems)

}
