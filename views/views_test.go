package views_test

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
	"text/template"

	"github.com/chris-weir/chrisweir/views"
)

func TestViewExecute(t *testing.T) {
	tpl := template.Must(template.New("test").Parse("<html>{{.}}</html>"))
	v := views.View{Template: tpl}

	tests := []struct {
		name         string
		data         interface{}
		expectedBody string
		expectedCode int
	}{
		{
			name:         "successful render",
			data:         "test data",
			expectedBody: "<html>test data</html>",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			v.Execute(rr, req, tc.data)

			if rr.Code != tc.expectedCode {
				t.Errorf("expected status code %d, got %d", tc.expectedCode, rr.Code)
			}

			if rr.Body.String() != tc.expectedBody {
				t.Errorf("expected body %q, got %q", tc.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestParseFromFile(t *testing.T) {
	mockFS := fstest.MapFS{
		"template.html": &fstest.MapFile{Data: []byte("<html>{{.}}</html>")},
	}
	tests := []struct {
		name      string
		fs        fs.FS
		patterns  []string
		wantError bool
	}{
		{
			name:      "successful parse",
			fs:        mockFS,
			patterns:  []string{"template.html"},
			wantError: false,
		},
		{
			name: "template parse error",
			fs:   mockFS, patterns: []string{"nonexistent.html"},
			wantError: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := views.ParseFromFile(tc.fs, tc.patterns...)
			if (err != nil) != tc.wantError {
				t.Errorf("expected error: %v, got: %v", tc.wantError, err)
			}
		})
	}
}
