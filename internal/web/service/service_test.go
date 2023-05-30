package service_test

import (
	"restaurant/internal/web/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopMenuItems(t *testing.T) {
	serviceStruct := service.ServiceStruct{}

	topMenuItems, err := serviceStruct.GetTopMenuItems("logfile_test1")

	assert.Equal(t, []string{"111", "123", "456"}, topMenuItems)
	assert.Nil(t, err)

	topMenuItems, err = serviceStruct.GetTopMenuItems("logfile_test2")

	assert.Equal(t, []string{"123", "456", "788"}, topMenuItems) //If more than 3 top menu items, 3 menu items with lease menu item id are selected
	assert.Nil(t, err)

	topMenuItems, err = serviceStruct.GetTopMenuItems("logfile_test")

	errorMessage := "open logfile_test: The system cannot find the file specified."
	assert.Equal(t, errorMessage, err.Error())

	topMenuItems, err = serviceStruct.GetTopMenuItems("logfile_test3")

	errorMessage = "Food menu id repeated for the given eater"
	assert.Equal(t, errorMessage, err.Error())
}
