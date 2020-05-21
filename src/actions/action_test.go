package actions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	libapi "github.com/iqdavidh/libapigo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func fnTestRequestGET(t2 *testing.T, a *assert.Assertions, queryParams string, group gin.HandlerFunc, codeRespuesta int) libapi.DicResp {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/test"
	r.GET(url, group)

	req, errReq := http.NewRequest(http.MethodGet, url+queryParams, nil)
	if errReq != nil {
		fmt.Println(errReq)
		t2.Fatalf("Couldn't create request: %v\n", errReq)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	a.True(w.Code == codeRespuesta, "No es el codigo Esperado")

	respuesta, errorDecode := libapi.DecodeBodyResponse(w.Body)

	a.True(errorDecode == nil, "Esperamos error nil "+fmt.Sprint(errorDecode))

	if errorDecode != nil {
		t2.Fatalf(fmt.Sprint(errorDecode))
	}

	return respuesta

}

func fnTestRequestPOST(t2 *testing.T, a *assert.Assertions, queryParams string, body string, handlerRequest gin.HandlerFunc, codeRespuesta int, dicHeader map[string]string) libapi.DicResp {
	gin.SetMode(gin.TestMode)
	ro := gin.Default()
	url := "/test"
	ro.POST(url, handlerRequest)

	req, errReq := http.NewRequest(http.MethodPost, url+queryParams, strings.NewReader(body))
	if errReq != nil {
		fmt.Println(errReq)
		t2.Fatalf("Couldn't create request: %v\n", errReq)
	}

	for k := range dicHeader {
		req.Header.Set(k, dicHeader[k])
	}

	w := httptest.NewRecorder()
	ro.ServeHTTP(w, req)

	a.True(w.Code == codeRespuesta, "No es el codigo Esperado")

	respuesta, errorDecode := libapi.DecodeBodyResponse(w.Body)

	a.True(errorDecode == nil, "Esperamos error nil "+fmt.Sprint(errorDecode))

	if errorDecode != nil {
		t2.Fatalf(fmt.Sprint(errorDecode))
	}

	return respuesta

}

func TestPruebaGet(t *testing.T) {

	a := assert.New(t)
	dicRespuesta := fnTestRequestGET(t, a, "", PruebaGet, 200)
	a.True(dicRespuesta.Success, "Esperabamos success")

}

func TestPruebaGetConQueryParams(t *testing.T) {
	a := assert.New(t)

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	url := "/pruebaget_con_queryparams"
	r.GET(url, PruebaGetConQueryParams)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, url+"?param1=valor_param1&param2=valor_param2", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	respuesta, errorDecode := libapi.DecodeBodyResponse(w.Body)

	a.True(errorDecode == nil, "Esperamos error nil")
	if errorDecode == nil {
		a.True(respuesta.Success, "No es success")
	}

}

func TestPruebaPost(t *testing.T) {

	a := assert.New(t)
	body := `{"user":"u","password":"p"}`
	dicHeader := map[string]string{"Token": "valort"}
	dicRespuesta := fnTestRequestPOST(t, a, "", body, PruebaPost, 200, dicHeader)
	a.True(dicRespuesta.Success, "Esperabamos success")

}
