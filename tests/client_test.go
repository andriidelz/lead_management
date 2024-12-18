// tests/client_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"lead_management/routes"

	"github.com/gin-gonic/gin"
)

func TestClientEndpoints(t *testing.T) {
	router := gin.Default()
	routes.RegisterClientRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clients", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}
}
