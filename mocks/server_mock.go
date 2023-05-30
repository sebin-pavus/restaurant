package mocks

import "github.com/gin-gonic/gin"

type MockRouter struct {
	*gin.Engine
}

func (m *MockRouter) GET(relativePath string, handlers ...gin.HandlerFunc) {
	m.Handle("GET", relativePath, handlers...)
}
