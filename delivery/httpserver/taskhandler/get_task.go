package taskhandler

import (
	"graph/param"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Hadler) GetTaskByID(c *gin.Context){
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req := param.GetTaskRequest{ID: uint(idInt)}

	resp, err := h.taskSvc.GetTaskByID(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}