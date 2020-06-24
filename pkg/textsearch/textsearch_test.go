package textsearch

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResponseBody(t *testing.T) {
	response := "test response body"

	// start test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(response))
	}))
	//

	b, _ := get(server.URL)
	if string(b) != response {
		t.Errorf("get() return bad response body, get: %s, expected: %s", b, []byte(response))
	}
}

func TestCountStringByURL(t *testing.T) {
	t.Run(`bad url case`, func(t *testing.T) {
		_, err := CountStringByURL("go", "bad_url")
		if err == nil {
			t.Errorf("Get no error while it's expected for bad url.")
		}
	})

	response := "Hello go!"

	// start test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(response))
	}))
	//

	t.Run(`count word "go"`, func(t *testing.T) {
		count, _ := CountStringByURL("go", server.URL)
		if count != 1 {
			t.Errorf("Expecting to get 1, get: %v", count)
		}
	})
	t.Run("count empty word", func(t *testing.T) {
		count, _ := CountStringByURL("", server.URL)
		if count != 10 {
			t.Errorf("Expecting to get 10, get: %v", count)
		}
	})
	t.Run("count phrase", func(t *testing.T) {
		count, _ := CountStringByURL("Hello go", server.URL)
		if count != 1 {
			t.Errorf("Expecting to get 1, get: %v", count)
		}
	})
}
