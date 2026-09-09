package taskhandler

import "github.com/gin-gonic/gin"

func (h Handler) SetRoutes(c *gin.Engine) {
	taskGroup := c.Group("/task")

	taskGroup.GET("/", h.GetTasks)
	taskGroup.POST("/", h.AddTask)
	taskGroup.GET("/:id", h.GetTaskByID)
}