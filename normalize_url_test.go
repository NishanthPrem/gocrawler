package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with query parameters",
			inputURL: "https://example.com/path?query=123",
			expected: "example.com/path?query=123",
		},
		{
			name:     "URL with fragment",
			inputURL: "https://example.com/path#section",
			expected: "example.com/path#section",
		},
		{
			name:     "URL with user info",
			inputURL: "https://user:password@example.com/path",
			expected: "user:password@example.com/path",
		},
		{
			name:     "URL with IP address",
			inputURL: "http://192.168.0.1/path",
			expected: "192.168.0.1/path",
		},
		{
			name:     "URL with IPv6 address",
			inputURL: "https://[2001:db8::1]:8080/path",
			expected: "[2001:db8::1]:8080/path",
		},
		{
			name:     "URL with no authority (relative path)",
			inputURL: "https:///path/to/resource",
			expected: "/path/to/resource",
		},
		{
			name:     "URL with multiple slashes",
			inputURL: "https://example.com//multiple/slashes",
			expected: "example.com//multiple/slashes",
		},
		{
			name:     "URL with percent encoding",
			inputURL: "https://example.com/path%20with%20spaces",
			expected: "example.com/path%20with%20spaces",
		},
		{
			name:     "URL with trailing slash",
			inputURL: "https://example.com/path/",
			expected: "example.com/path/",
		},
		{
			name:     "URL with special characters in path",
			inputURL: "https://example.com/path/@special#chars",
			expected: "example.com/path/@special#chars",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
