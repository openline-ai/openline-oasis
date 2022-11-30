package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiRouter *gin.Engine

func init() {
	apiRouter = gin.Default()
	route := apiRouter.Group("/")

	route.Use(apiKeyChecker(TEST_API_KEY))
	addHealthRoutes(route)
}

const TEST_API_KEY = "2e0f7ee7-f874-4124-a805-b97725ec5fab"

func TestApiKeyOK(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	req.Header.Set("x-openline-api-key", TEST_API_KEY)
	apiRouter.ServeHTTP(w, req)
	if !assert.Equal(t, w.Code, 200) {
		return
	}
	var resp HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("TestHealthCheck, unable to decode json: %v\n", err.Error())
		return
	}
	assert.Equal(t, resp.Status, "ok")
}

func TestApiKeyNOK(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	req.Header.Set("x-openline-api-key", TEST_API_KEY+"-nok")
	apiRouter.ServeHTTP(w, req)
	if !assert.Equal(t, w.Code, 403) {
		return
	}
}

func TestApiKeyMissing(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	apiRouter.ServeHTTP(w, req)
	if !assert.Equal(t, w.Code, 403) {
		return
	}
}
