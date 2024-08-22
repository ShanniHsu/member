package main

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"member/models"
	router2 "member/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := router2.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPostUser(t *testing.T) {
	router := router2.SetupRouter()
	router = router2.PostUser(router)

	w := httptest.NewRecorder()

	// Create an example user for testing
	exampleUser := models.User{
		Account: "shanni",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Compare the response body with the json data of exampleUser
	assert.Equal(t, string(userJson), w.Body.String())
}
