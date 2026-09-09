package taskhandler

import "github.com/gin-gonic/gin"

func (h Hadler) SetRoutes(c *gin.Engine) {
	taskGroup := c.Group("/task")

	taskGroup.POST("/", h.AddTask)
	taskGroup.GET("/:id", h.GetTaskByID)
}