package service

import (
	"go-napi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *CommandHandler) Register(c *gin.Context) {

	a := model.User{}

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}
