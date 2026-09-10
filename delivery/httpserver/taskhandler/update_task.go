package taskhandler

import (
	"graph/param"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateTask godoc
// @Summary      Update a task
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        id      path  int                    true "Task ID"
// @Param        request body  param.UpdateTaskRequest true "Task data"
// @Success      200 {object} map[string]param.UpdateTaskResponse
// @Failure      400 {object} map[string]string
// @Router       /task/{id} [put]
func (h Handler) UpdateTask(c *gin.Context) {
	var req param.UpdateTaskRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.ID = uint(idInt)

	resp, err := h.taskSvc.UpdateTask(c.Request.Context(), req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}