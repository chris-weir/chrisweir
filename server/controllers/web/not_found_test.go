package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controllers "github.com/chris-weir/chrisweir/server/controllers/web"
)

func TestNotFoundReturns(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	controllers.NotFound(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status: %v got: %v", http.StatusNotFound, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read request data %s", err)
	}

	mainHeaderExists := strings.Contains(string(data), "Page not found &#128546")
	if !mainHeaderExists {
		t.Errorf("Head missing from data. output: %s", string(data))
	}
}
