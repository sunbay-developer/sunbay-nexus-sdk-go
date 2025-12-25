package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

const (
	sdkName        = "SunbayNexusSDK-Go"
	language       = "Go"
	unknownVersion = "unknown"
	defaultSDKVersion = "1.0.0"
)

var (
	cachedUserAgent string
	once            sync.Once
)

// GetUserAgent generates User-Agent string
// Format: SunbayNexusSDK-Go/{SDKVersion} Go/{GoVersion} {OS}/{OSVersion}
func GetUserAgent() string {
	once.Do(func() {
		cachedUserAgent = buildUserAgent()
	})
	return cachedUserAgent
}

// buildUserAgent builds User-Agent string
func buildUserAgent() string {
	sdkVersion := getSDKVersion()
	goVersion := runtime.Version()
	osName := normalizeOSName(runtime.GOOS)
	osVersion := getOSVersion()

	return fmt.Sprintf("%s/%s %s/%s %s/%s",
		sdkName, sdkVersion,
		language, goVersion,
		osName, osVersion)
}

// getSDKVersion gets SDK version
// First tries to get from build info, then from build-time injected variable, finally uses default
func getSDKVersion() string {
	// Try to get from build info (runtime/debug)
	if info, ok := debug.ReadBuildInfo(); ok {
		// Check if there's a version in build settings
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" || setting.Key == "vcs.version" {
				if setting.Value != "" && setting.Value != "(devel)" {
					return setting.Value
				}
			}
		}
		// Use module version if available
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	// Try build-time injected variable (set via -ldflags)
	// This would be set at build time: go build -ldflags "-X github.com/sunbay-developer/sunbay-nexus-sdk-go/util.sdkVersion=1.0.0"
	// For now, we'll use a default version
	return defaultSDKVersion
}

// normalizeOSName normalizes OS name
func normalizeOSName(osName string) string {
	switch osName {
	case "darwin":
		return "Darwin"
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "freebsd", "openbsd", "netbsd":
		return "Unix"
	default:
		if osName == "" {
			return unknownVersion
		}
		return osName
	}
}

// getOSVersion gets OS version using standard library only
func getOSVersion() string {
	switch runtime.GOOS {
	case "darwin":
		// macOS: use sw_vers command
		return getDarwinVersion()
	case "linux":
		// Linux: try to read /etc/os-release
		return getLinuxVersion()
	case "windows":
		// Windows: can use ver command, but requires exec
		return getWindowsVersion()
	default:
		return unknownVersion
	}
}

// getDarwinVersion gets macOS version
func getDarwinVersion() string {
	cmd := exec.Command("sw_vers", "-productVersion")
	output, err := cmd.Output()
	if err != nil {
		return unknownVersion
	}
	return strings.TrimSpace(string(output))
}

// getLinuxVersion gets Linux version
func getLinuxVersion() string {
	// Try to read /etc/os-release
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return unknownVersion
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "VERSION_ID=") {
			version := strings.TrimPrefix(line, "VERSION_ID=")
			version = strings.Trim(version, `"`)
			return version
		}
	}

	// Fallback: try /etc/redhat-release or /etc/debian_version
	if data, err := os.ReadFile("/etc/redhat-release"); err == nil {
		return strings.TrimSpace(string(data))
	}
	if data, err := os.ReadFile("/etc/debian_version"); err == nil {
		return strings.TrimSpace(string(data))
	}

	return unknownVersion
}

// getWindowsVersion gets Windows version
func getWindowsVersion() string {
	cmd := exec.Command("cmd", "/c", "ver")
	output, err := cmd.Output()
	if err != nil {
		return unknownVersion
	}
	version := strings.TrimSpace(string(output))
	// Extract version number from "Microsoft Windows [Version 10.0.xxxxx]"
	if idx := strings.Index(version, "Version"); idx != -1 {
		version = version[idx+8:]
		version = strings.Trim(version, "[]")
		return version
	}
	return version
}

