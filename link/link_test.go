package link

import (
	"net/url"
	"testing"
)

func TestIsUrl(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"https://www.google.com", true},
		{"https://www.example.com", true},
		{"ftp://ftp.example.com", true},
		{"justastring", false},
		{"https://", false},
		{"www.google.com", false},
	}

	for _, tc := range testCases {
		result := IsUrl(tc.input)
		if result != tc.expected {
			t.Errorf("IsUrl(%s) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestGetHostname(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"www.google.com", "google.com"},
		{"example.com", "example.com"},
		{"www.instagram.com", "instagram.com"},
		{"subdomain.example.com", "subdomain.example.com"},
	}

	for _, tc := range testCases {
		result := getHostname(tc.input)
		if result != tc.expected {
			t.Errorf("getHostname(%s) = %s; expected %s", tc.input, result, tc.expected)
		}
	}
}

func TestIsFixableUrl(t *testing.T) {
	testCases := []struct {
		inputURL  string
		isFixable bool
	}{
		{"https://instagram.com/photo", true},
		{"https://twitter.com/status", true},
		{"https://x.com/content", true},
		{"https://reddit.com/r/sub", true},
		{"https://tiktok.com/video", true},
		{"https://zietlow.io", false},
	}

	for _, tc := range testCases {
		parsedURL, _ := url.Parse(tc.inputURL)
		link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}
		result := link.IsFixableUrl()
		if result != tc.isFixable {
			t.Errorf("IsFixableUrl() for URL %s = %v; expected %v", tc.inputURL, result, tc.isFixable)
		}
	}
}

// TestFix function tests the Fix method of Link struct
func TestFix(t *testing.T) {
	testCases := []struct {
		inputURL string
		expected string
	}{
		{"https://instagram.com/photo", "https://ddinstagram.com/photo"},
		{"https://twitter.com/status", "https://fxtwitter.com/status"},
		{"https://x.com/content", "https://fixvx.com/content"},
		{"https://reddit.com/r/sub", "https://vxreddit.com/r/sub"},
		{"https://tiktok.com/video", "https://vxtiktok.com/video"},
	}

	for _, tc := range testCases {
		parsedURL, _ := url.Parse(tc.inputURL)
		link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}
		result := link.Fix()
		if result != tc.expected {
			t.Errorf("Fix() for URL %s = %s; expected %s", tc.inputURL, result, tc.expected)
		}
	}
}

// Additional individual tests for each fix function can be structured in a similar manner
func TestFixInstagram(t *testing.T) {
	testURL := "https://instagram.com/photo"
	expected := "https://ddinstagram.com/photo"

	parsedURL, _ := url.Parse(testURL)
	link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}

	result := fixInstagram(link)
	if result != expected {
		t.Errorf("fixInstagram() = %s; expected %s", result, expected)
	}
}

// TestFixTwitter tests the fixTwitter function
func TestFixTwitter(t *testing.T) {
	testURL := "https://twitter.com/status"
	expected := "https://fxtwitter.com/status"

	parsedURL, _ := url.Parse(testURL)
	link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}

	result := fixTwitter(link)
	if result != expected {
		t.Errorf("fixTwitter() = %s; expected %s", result, expected)
	}
}

// TestFixX tests the fixX function
func TestFixX(t *testing.T) {
	testURL := "https://x.com/content"
	expected := "https://fixvx.com/content"

	parsedURL, _ := url.Parse(testURL)
	link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}

	result := fixX(link)
	if result != expected {
		t.Errorf("fixX() = %s; expected %s", result, expected)
	}
}

// TestFixReddit tests the fixReddit function
func TestFixReddit(t *testing.T) {
	testURL := "https://reddit.com/r/sub"
	expected := "https://vxreddit.com/r/sub"

	parsedURL, _ := url.Parse(testURL)
	link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}

	result := fixReddit(link)
	if result != expected {
		t.Errorf("fixReddit() = %s; expected %s", result, expected)
	}
}

// TestFixTikTok tests the fixTikTok function
func TestFixTikTok(t *testing.T) {
	testURL := "https://tiktok.com/video"
	expected := "https://vxtiktok.com/video"

	parsedURL, _ := url.Parse(testURL)
	link := &Link{URL: parsedURL, Hostname: parsedURL.Hostname()}

	result := fixTikTok(link)
	if result != expected {
		t.Errorf("fixTikTok() = %s; expected %s", result, expected)
	}
}
