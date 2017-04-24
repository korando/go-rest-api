package main

import (
	apilib "./lib"
	users "./users"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "ducvu:1234567980@tcp(mydbinstance.cywu6qvxsqy7.ap-southeast-1.rds.amazonaws.com:3306)/godb")
	apilib.CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	users.Dbmap = dbmap

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
