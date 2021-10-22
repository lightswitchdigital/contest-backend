package server

import (
	"github.com/gin-gonic/gin"
	"time"
)

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

func (h *Handler) HandleGetAvailableTransport(c *gin.Context) {
	c.JSON(200, []gin.H{
		gin.H{
			"id":             987234,
			"name":           "Автобус 41",
			"type":           "bus",
			"occupancy":      "high",
			"time_to_arrive": time.Now().Add(time.Minute * 2).Unix(),
		},
		gin.H{
			"id":             4287622,
			"name":           "Автобус 70",
			"type":           "bus",
			"occupancy":      "average",
			"time_to_arrive": time.Now().Add(time.Minute * 10).Unix(),
		},
		gin.H{
			"id":             87623483,
			"name":           "Трамвай 3",
			"type":           "tramway",
			"occupancy":      "low",
			"time_to_arrive": time.Now().Add(time.Minute * 5).Unix(),
		},
		gin.H{
			"id":             77234982,
			"name":           "Троллейбус 15",
			"type":           "trolleybus",
			"occupancy":      "low",
			"time_to_arrive": time.Now().Add(time.Minute * 3).Unix(),
		},
	})
}
