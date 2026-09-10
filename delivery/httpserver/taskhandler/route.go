package taskhandler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

)

func (h Handler) SetRoutes(c *gin.Engine) {
	
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	taskGroup := c.Group("/task")

	taskGroup.GET("/", h.GetTasks)
	taskGroup.POST("/", h.AddTask)
	taskGroup.GET("/:id", h.GetTaskByID)
	taskGroup.PUT("/:id", h.UpdateTask)
	taskGroup.DELETE("/:id", h.DeleteTask)
}