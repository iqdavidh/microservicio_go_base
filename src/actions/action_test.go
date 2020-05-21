package actions

import (
	libapi "github.com/iqdavidh/libapigo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPruebaGet(t *testing.T) {
	a := assert.New(t)
	configTest := libapi.FactoryConfigTestBasic(nil)
	dicRespuesta := libapi.TestBasicRequestGET(t, a, PruebaGet, configTest)
	a.True(dicRespuesta.Success, "Esperabamos success")
}

func TestPruebaGetConQueryParams(t *testing.T) {
	a := assert.New(t)

	configTest := libapi.FactoryConfigTestBasic(nil)
	configTest.QueryParams = "param1=valor_param1&param2=valor_param2"

	dicRespuesta := libapi.TestBasicRequestGET(t, a, PruebaGetConQueryParams, configTest)
	a.True(dicRespuesta.Success, "Esperabamos success")

	//Validacion especifica -----------
	a.Equal("valor_param1", dicRespuesta.Data["param1"])
	a.Equal("valor_param2", dicRespuesta.Data["param2"])
}

func TestPruebaGetConUrlParams(t *testing.T) {
	a := assert.New(t)

	configTest := libapi.FactoryConfigTestBasic(nil)
	configTest.UrlParamsValor = "/bart"
	configTest.UrlParamsPatron = "/:nombre"

	dicRespuesta := libapi.TestBasicRequestGET(t, a, PruebaGetConUrlParams, configTest)
	a.True(dicRespuesta.Success, "Esperabamos success")

	//Validacion especifica -----------
	a.Equal("bart", dicRespuesta.Data["nombre"])
}

func TestPruebaPost(t *testing.T) {
	a := assert.New(t)

	configTest := libapi.FactoryConfigTestBasic(map[string]string{"Token": "valort"})
	configTest.Body = `{"user":"u","password":"p"}`

	dicRespuesta := libapi.TestBasicRequestPOST(t, a, PruebaPost, configTest)
	a.True(dicRespuesta.Success, "Esperabamos success")
}
