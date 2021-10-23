package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Handler struct {
	data map[int]gin.H
}

func NewHandler() *Handler {
	return &Handler{
		data: map[int]gin.H{
			987234: gin.H{
				"id":             987234,
				"name":           "Автобус 41",
				"type":           "bus",
				"occupancy":      "high",
				"time_to_arrive": time.Now().Add(time.Minute * 2).Unix(),
			},
			4287622: gin.H{
				"id":             4287622,
				"name":           "Автобус 70",
				"type":           "bus",
				"occupancy":      "average",
				"time_to_arrive": time.Now().Add(time.Minute * 10).Unix(),
			},
			87623483: gin.H{
				"id":             87623483,
				"name":           "Трамвай 3",
				"type":           "tramway",
				"occupancy":      "low",
				"time_to_arrive": time.Now().Add(time.Minute * 5).Unix(),
			},
			77234982: gin.H{
				"id":             77234982,
				"name":           "Троллейбус 15",
				"type":           "trolleybus",
				"occupancy":      "low",
				"time_to_arrive": time.Now().Add(time.Minute * 3).Unix(),
			},
		},
	}
}

func (h *Handler) HandleIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleGetAvailableTransport(c *gin.Context) {
	arr := make([]gin.H, 0, len(h.data))
	for _, el := range h.data {
		arr = append(arr, el)
	}
	c.JSON(200, arr)
}

func (h *Handler) FindHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("error parsing id param: %v", err)
	}

	c.JSON(200, gin.H{
		"result": h.data[id],
	})
}
