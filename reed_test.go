package reed_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"rohitsingh/reed"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// setupAPI is a helper function that sets up a httptest server
func setupAPI(t *testing.T, r *mux.Router) *httptest.Server {
	t.Helper() // Mark the function as test helper
	// Create a new Mux to handle incoming requests

	ts := httptest.NewServer(r)
	// Return the ts instance to allow test function
	// to attach custom handlers
	return ts
}

// TestReplyError uses table-driven testing to make sure
// the ReplyError function is working as expected
func TestReplyError(t *testing.T) {
	// Use table-driven testing to cover various edge cases
	testCases := []struct {
		name    string // Name of the test
		status  int    // HTTP response code
		content string // response body
	}{
		{"SimpleErrorMessage", http.StatusNotFound, "not lookin good"},
	}
	for _, tc := range testCases {
		t.Run("Test"+tc.name, func(t *testing.T) {
			// Define a function wrapper that will be used
			// to handle get requests to the root
			handler := func(w http.ResponseWriter, r *http.Request) {
				reed.ReplyError(w, r, tc.status, tc.content)
			}
			// Attach the handler to the root of the server
			r := mux.NewRouter()
			r.HandleFunc("/", handler)
			// Create a HTTP test server
			ts := setupAPI(t, r)
			defer ts.Close()
			// Send a GET request to the root
			resp, err := http.Get(ts.URL + "/")
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tc.status {
				t.Fatalf("Expected status code %s, instead got %s",
					http.StatusText(tc.status),
					http.StatusText(resp.StatusCode),
				)
			}
			// Check that we have the content that we expected
			var body []byte
			if body, err = io.ReadAll(resp.Body); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), tc.content) {
				t.Fatalf("Expected %q, got %q.", tc.content, string(body))
			}
		})
	}
}

// TestReplyTextContent uses table-driven testing to make sure
// the ReplyTextContent function is working as expected
func TestReplyTextContent(t *testing.T) {
	// Use table-driven testing to cover various edge cases
	testCases := []struct {
		name    string // Name of the test
		status  int    // HTTP response code
		content string // response body
	}{
		{"SimpleMessageBody", http.StatusOK, "lookin good"},
		{"MessageWithNewlines", http.StatusAccepted, "lookin\ngood\n"},
	}
	for _, tc := range testCases {
		t.Run("Test"+tc.name, func(t *testing.T) {
			// Define a function wrapper that will be used
			// to handle get requests to the root
			handler := func(w http.ResponseWriter, r *http.Request) {
				reed.ReplyTextContent(w, r, tc.status, tc.content)
			}
			// Attach the handler to the root of the server
			r := mux.NewRouter()
			r.HandleFunc("/", handler)
			// Create a HTTP test server
			ts := setupAPI(t, r)
			defer ts.Close()
			// Send a GET request to the root
			resp, err := http.Get(ts.URL + "/")
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tc.status {
				t.Fatalf("Expected status code %s, instead got %s",
					http.StatusText(tc.status),
					http.StatusText(resp.StatusCode),
				)
			}
			// Check that we have the content that we expected
			var body []byte
			if body, err = io.ReadAll(resp.Body); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), tc.content) {
				t.Fatalf("Expected %q, got %q.", tc.content, string(body))
			}
		})
	}
}
