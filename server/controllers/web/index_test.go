package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controllers "github.com/chris-weir/chrisweir/server/controllers/web"
)

func TestIndexReturnsView(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	controllers.Index(w, r)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read request data %s", err)
	}

	mainHeaderExists := strings.Contains(string(data), "Hi, I'm Chris &#128075")
	if !mainHeaderExists {
		t.Errorf("Head missing from data. output: %s", string(data))
	}
}
