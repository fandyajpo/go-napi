package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *CommandHandler) Book(c *gin.Context) {
	books := make(map[string]string)
	books["title"] = "mobile legends"
	books["title"] = "napi tester"
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (m *CommandHandler) StringReturn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hai"})
}

func (m *CommandHandler) NumberReturn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hai"})
}
