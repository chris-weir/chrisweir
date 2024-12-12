package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chris-weir/chrisweir/server"
)

func TestServerRuns(t *testing.T) {
	server := server.GetServer()
	server.MountHandlers()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	server.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code: %d. Got: %d\n", http.StatusOK, rr.Code)
	}
}
