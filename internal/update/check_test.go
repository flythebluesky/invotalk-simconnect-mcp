package update

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckLatest_SkipsDev(t *testing.T) {
	// Should return immediately without making any HTTP call.
	CheckLatest("dev")
}

func TestCheckLatest_SameVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(githubRelease{
			TagName: "v1.0.0",
			HTMLURL: "https://github.com/test/releases/tag/v1.0.0",
		})
	}))
	defer srv.Close()

	origURL := releaseURL
	setReleaseURL(srv.URL)
	defer setReleaseURL(origURL)

	// Same version — should not panic or error.
	CheckLatest("v1.0.0")
}

func TestCheckLatest_NewerVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(githubRelease{
			TagName: "v2.0.0",
			HTMLURL: "https://github.com/test/releases/tag/v2.0.0",
		})
	}))
	defer srv.Close()

	origURL := releaseURL
	setReleaseURL(srv.URL)
	defer setReleaseURL(origURL)

	// Different version — should log but not panic.
	CheckLatest("v1.0.0")
}

func TestCheckLatest_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	origURL := releaseURL
	setReleaseURL(srv.URL)
	defer setReleaseURL(origURL)

	// Server error — should handle gracefully.
	CheckLatest("v1.0.0")
}

func TestCheckLatest_NetworkError(t *testing.T) {
	origURL := releaseURL
	setReleaseURL("http://localhost:1") // nothing listening
	defer setReleaseURL(origURL)

	CheckLatest("v1.0.0")
}
