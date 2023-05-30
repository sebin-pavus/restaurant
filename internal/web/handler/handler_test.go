package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"restaurant/internal/model"
	"restaurant/internal/web/handler"
	"restaurant/mocks"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTopMenuItemsSucscess(t *testing.T) {
	router := mocks.MockRouter{Engine: gin.New()}

	mockHandler := handler.HandlerStruct{Service: mocks.MockServiceSuccess{}}

	router.GET("/top_items", mockHandler.GetTopMenuItems)

	req, err := http.NewRequest("GET", "/top_items", nil)
	assert.NoError(t, err)

	// Create a mock HTTP response
	resp := httptest.NewRecorder()

	// Dispatch the request using the mock router
	router.ServeHTTP(resp, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	var arr []string
	err = json.Unmarshal([]byte(resp.Body.String()), &arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.Equal(t, mocks.Sample, arr)
}

func TestTopMenuItemsFailure(t *testing.T) {
	router := mocks.MockRouter{Engine: gin.New()}

	mockHandler := handler.HandlerStruct{Service: mocks.MockServiceFailure{}}

	router.GET("/top_items", mockHandler.GetTopMenuItems)

	req, err := http.NewRequest("GET", "/top_items", nil)
	assert.NoError(t, err)

	// Create a mock HTTP response
	resp := httptest.NewRecorder()

	// Dispatch the request using the mock router
	router.ServeHTTP(resp, req)

	unquotedStr, err := strconv.Unquote(resp.Body.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Assert the response status code
	errResponse := model.ErrorResponse{ErrorMessage: "Error getting menu items"}
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, errResponse, unquotedStr)
}
