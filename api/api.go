package api

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/karsteneugene/absen-backend/api/v1"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Acess-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Api() {
	router := gin.Default()

	router.Use(cors())

	api := router.Group("/api")

	v1 := api.Group("/v1")

	v1.GET("/users", apiv1.GetUsers)
	v1.GET("/users/:id", apiv1.GetUserByID)
	v1.POST("/users", apiv1.PostUser)
	v1.PUT("/users/:id", apiv1.UpdateUser)
	v1.DELETE("/users/:id", apiv1.DeleteUser)

	router.Run()
}
