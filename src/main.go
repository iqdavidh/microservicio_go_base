package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iqdavidh/libapi"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/pruebaget", func(ctx *gin.Context) {
		data := libapi.DicJson{"ope": "pruebaget sin parametros"}
		libapi.Success(ctx, data)
	})

	r.GET("/pruebaget_con_queryparams", func(ctx *gin.Context) {
		dicParams := libapi.GetDataCleanFromQP(ctx, []string{"param1", "param2"})
		libapi.Success(ctx, dicParams)
	})

	// Get user value
	r.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		value, ok := db[user]
		if ok {
			//ctx.JSON(http.StatusOK, libapi.DicJson{"user": user, "value": value})
			ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

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
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
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
