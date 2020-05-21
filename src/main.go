package main

import (
	"curso/microservice/src/actions"
	"fmt"
	"github.com/gin-gonic/gin"
	libapi "github.com/iqdavidh/libapigo"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/pruebaget", actions.PruebaGet)

	r.GET("/pruebaget_con_queryparams", actions.PruebaGetConQueryParams)

	r.GET("/pruebaget_conparams/:nombre", actions.PruebaGetConUrlParams)

	r.GET("/pruebaget_error", func(ctx *gin.Context) {
		libapi.RespuestaError(ctx, "algo porque si")
	})

	r.POST("/pruebapost", actions.PruebaPost)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	fmt.Println("\n\n***************************************\nRuning server 4 ***********************\n***************************************")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
