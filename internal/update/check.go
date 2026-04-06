package update

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

var releaseURL = "https://api.github.com/repos/flythebluesky/invotalk-simconnect-mcp/releases/latest"

func setReleaseURL(url string) { releaseURL = url }

type githubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

// CheckLatest compares the running version against the latest GitHub release.
// Logs a message if a newer version is available. Designed to run in a goroutine.
func CheckLatest(currentVersion string) {
	if currentVersion == "dev" {
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(releaseURL)
	if err != nil {
		slog.Debug("update check failed", "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Debug("update check: unexpected status", "status", resp.StatusCode)
		return
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		slog.Debug("update check: decode error", "error", err)
		return
	}

	latest := strings.TrimPrefix(release.TagName, "v")
	current := strings.TrimPrefix(currentVersion, "v")

	if latest != current {
		slog.Info("new version available",
			"current", currentVersion,
			"latest", release.TagName,
			"download", release.HTMLURL,
		)
	}
}
