package taskhandler

import (
	"graph/param"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTask godoc
// @Summary      Add a new task
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        request body param.AddTaskRequest true "Task data"
// @Success      201 {object} map[string]param.AddTaskResponse
// @Failure      400 {object} map[string]string
// @Router       /task/ [post]
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