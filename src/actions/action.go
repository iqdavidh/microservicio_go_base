package actions

import (
	"github.com/gin-gonic/gin"
	libapi "github.com/iqdavidh/libapigo"
)

func PruebaGet(ctx *gin.Context) {
	data := libapi.DicJson{"ope": "pruebaget sin parametros vertsion2"}
	libapi.RespuestaSuccess(ctx, data)
}

func PruebaGetConQueryParams(ctx *gin.Context) {
	dicParams := libapi.GetDataCleanFromQP(ctx, []string{"param1", "param2"})
	libapi.RespuestaSuccess(ctx, dicParams)
}

func PruebaGetConUrlParams(ctx *gin.Context) {
	nombre := ctx.Params.ByName("nombre")
	libapi.RespuestaSuccess(ctx, libapi.DicJson{"nombre": nombre})
}

type PruebaPostData struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func PruebaPost(ctx *gin.Context) {
	var dataPost PruebaPostData

	err := ctx.ShouldBindJSON(&dataPost)
	if err != nil {
		libapi.RespuestaError(ctx, err.Error())
		return
	}

	token := ctx.GetHeader("Token")

	respuesta := libapi.DicJson{
		"user":         dataPost.User,
		"password":     dataPost.Password,
		"header token": token,
		"version":      1,
	}

	libapi.RespuestaSuccess(ctx, respuesta)
}
