package taskhandler

import (
	"graph/param"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) AddTask(c *gin.Context) {
	var req param.AddTaskRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.taskSvc.AddTask(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": resp,
	})
}