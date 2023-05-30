package handler

import (
	"net/http"
	"os"

	"restaurant/internal/model"
	"restaurant/internal/web/service"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetTopMenuItems(c *gin.Context)
}

type HandlerStruct struct {
	Service service.ServiceInterface
}

func (h HandlerStruct) GetTopMenuItems(c *gin.Context) {
	logFile := os.Getenv("LOG_FILE")

	topMenuItems, err := h.Service.GetTopMenuItems(logFile)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	c.IndentedJSON(http.StatusOK, topMenuItems)

}
