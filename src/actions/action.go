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
