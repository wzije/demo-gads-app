package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userCtrl = NewUserController()

func Test_userController_Ping(t *testing.T) {

	router := gin.Default()
	router.GET("/ping", userCtrl.Ping)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func Test_userController_Fetch(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", userCtrl.Fetch)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
