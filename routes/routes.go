package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/punit1997/ApiResponse/controller"
)

func InitRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	route.POST("/eecom_job_order", controller.JobOrder)

	return route
}
