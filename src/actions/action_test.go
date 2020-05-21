package actions

import (
	"github.com/gin-gonic/gin"
	libapi "github.com/iqdavidh/libapigo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPruebaGet(t *testing.T) {

	a := assert.New(t)
	dicRespuesta := libapi.TestBasicRequestGET(t, a, "", PruebaGet, 200, nil)
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
	dicRespuesta := libapi.TestBasicRequestPOST(t, a, "", body, PruebaPost, 200, dicHeader)
	a.True(dicRespuesta.Success, "Esperabamos success")

}
