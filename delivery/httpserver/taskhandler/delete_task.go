package taskhandler

import (
	"graph/param"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteTask(c *gin.Context) {
	var req param.DeleteTaskRequest
	
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = uint(idInt)

	err = h.taskSvc.DeleteTask(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}