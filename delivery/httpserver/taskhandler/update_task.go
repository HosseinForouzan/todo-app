package taskhandler

import (
	"graph/param"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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