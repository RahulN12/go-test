package tests

import (
	"bytes"
	"encoding/json"
	"github.com/RahulN12/go-test/controllers"
	"github.com/RahulN12/go-test/repository"
	"github.com/RahulN12/go-test/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	repo := repository.NewGraphRepository()
	service := service.NewGraphService(repo)
	controller := controllers.NewGraphController(service)

	r.POST("/graph", controller.PostGraph)
	r.GET("/graph/:id/shortest-path", controller.GetShortestPath)
	r.DELETE("/graph/:id", controller.DeleteGraph)

	return r
}

func TestPostGraph(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B"],"edges":[["A","B"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"id":"1"}`, resp.Body.String())
}

func TestGetShortestPath(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B"],"edges":[["A","B"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	id := result["id"]

	req, _ = http.NewRequest("GET", "/graph/"+id+"/shortest-path?start=A&end=B", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetShortestPathSinglePath(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B"],"edges":[["A","B"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	id := result["id"]

	req, _ = http.NewRequest("GET", "/graph/"+id+"/shortest-path?start=A&end=B", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"path":["A","B"]}`, resp.Body.String())
}

func TestGetShortestPathMultiplePath(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B","C","D","E"],"edges":[["A","B"],["B","C"],["C","D"],["D","E"],["B","D"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	id := result["id"]

	req, _ = http.NewRequest("GET", "/graph/"+id+"/shortest-path?start=A&end=D", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"path":["A","B","D"]}`, resp.Body.String())
}

func TestGetShortestPathSameStartEnd(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B"],"edges":[["A","B"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	id := result["id"]

	req, _ = http.NewRequest("GET", "/graph/"+id+"/shortest-path?start=A&end=A", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteGraph(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	graphData := `{"nodes":["A","B"],"edges":[["A","B"]]}`
	req, _ := http.NewRequest("POST", "/graph", bytes.NewBufferString(graphData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	id := result["id"]

	req, _ = http.NewRequest("DELETE", "/graph/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest("GET", "/graph/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}
