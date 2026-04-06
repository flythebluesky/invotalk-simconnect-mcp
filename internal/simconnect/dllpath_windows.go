//go:build windows

package simconnect

import (
	"log/slog"
	"os"
	"path/filepath"
)

const dllName = "SimConnect.dll"

// findDLL searches for SimConnect.dll in well-known locations.
// If dllPath is non-empty, it is returned as-is (explicit override).
func findDLL(dllPath string) string {
	if dllPath != "" {
		slog.Debug("using explicit SimConnect DLL path", "path", dllPath)
		return dllPath
	}

	candidates := buildSearchPaths()
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			slog.Info("found SimConnect.dll", "path", p)
			return p
		}
	}

	slog.Debug("SimConnect.dll not found in search paths, falling back to default DLL search order")
	return dllName
}

func buildSearchPaths() []string {
	var paths []string

	// Next to the executable.
	if exe, err := os.Executable(); err == nil {
		paths = append(paths, filepath.Join(filepath.Dir(exe), dllName))
	}

	// MSFS 2024 SDK env var.
	if sdk := os.Getenv("MSFS2024_SDK"); sdk != "" {
		paths = append(paths, filepath.Join(sdk, "SimConnect SDK", "lib", dllName))
	}

	// MSFS 2020 SDK env var.
	if sdk := os.Getenv("MSFS_SDK"); sdk != "" {
		paths = append(paths, filepath.Join(sdk, "SimConnect SDK", "lib", dllName))
	}

	// Hardcoded default SDK install paths.
	paths = append(paths,
		filepath.Join(`C:\MSFS 2024 SDK`, "SimConnect SDK", "lib", dllName),
		filepath.Join(`C:\MSFS SDK`, "SimConnect SDK", "lib", dllName),
	)

	return paths
}
