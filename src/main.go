package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iqdavidh/libapi"
	"net/http"
)

type PruebaPostData struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

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

	r.GET("/pruebaget_conparams/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Params.ByName("nombre")
		libapi.Success(ctx, libapi.DicJson{"nombre": nombre})
	})

	r.POST("/pruebapost", func(ctx *gin.Context) {

		var dataPost PruebaPostData

		err := ctx.ShouldBindJSON(&dataPost)
		if err != nil {
			libapi.Error(ctx, err.Error())
			return
		}

		respuesta := libapi.DicJson{
			"user":     dataPost.User,
			"password": dataPost.Password,
		}

		libapi.Success(ctx, respuesta)
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
