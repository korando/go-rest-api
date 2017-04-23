package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	users "./users"
)




func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.GET("/users", users.GetUsers)
		v1.GET("/users/:id", users.GetUser)
		v1.POST("/users", users.PostUser)
		v1.PUT("/users/:id", users.UpdateUser)
		v1.DELETE("/users/:id", users.DeleteUser)
	}

	r.Run(":8080")
}

