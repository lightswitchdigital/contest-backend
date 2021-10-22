package server

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
