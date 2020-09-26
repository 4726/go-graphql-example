package web

import (
	"log"

	api "github.com/4726/go-graphql-example/graphql"
	"github.com/gin-gonic/gin"
)

func RunServer() error {
	r := gin.Default()

	r.GET("/graphql", func(c *gin.Context) {
		res := api.Do(c.Query("query"))
		if len(res.Errors) > 0 {
			log.Print(res.Errors)
			c.JSON(500, "internal server error")
		} else {
			c.JSON(200, res)
		}
	})

	return r.Run(":12345")
}
