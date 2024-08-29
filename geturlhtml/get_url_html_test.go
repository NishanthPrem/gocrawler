package gourlhtml

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		expectErr bool
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
                <html>
                    <body>
                        <a href="/path/one">
                            <span>Boot.dev</span>
                        </a>
                        <a href="https://other.com/path/one">
                            <span>Boot.dev</span>
                        </a>
                    </body>
                </html>`,
			expected:  []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			expectErr: false,
		},
		{
			name:     "no URLs in HTML",
			inputURL: "https://example.com",
			inputBody: `
                <html>
                    <body>
                        <p>No links here!</p>
                    </body>
                </html>`,
			expected:  []string{},
			expectErr: false,
		},
		{
			name:      "empty HTML body",
			inputURL:  "https://example.com",
			inputBody: "",
			expected:  []string{},
			expectErr: false,
		},
		{
			name:     "relative URL only",
			inputURL: "https://example.com",
			inputBody: `
                <html>
                    <body>
                        <a href="/relative/path">
                            <span>Relative Link</span>
                        </a>
                    </body>
                </html>`,
			expected:  []string{"https://example.com/relative/path"},
			expectErr: false,
		},
		{
			name:     "multiple relative URLs",
			inputURL: "https://example.com",
			inputBody: `
                <html>
                    <body>
                        <a href="/path/one">Link One</a>
                        <a href="/path/two">Link Two</a>
                    </body>
                </html>`,
			expected: []string{
				"https://example.com/path/one",
				"https://example.com/path/two",
			},
			expectErr: false,
		},
		{
			name:     "invalid URL in HTML",
			inputURL: "https://example.com",
			inputBody: `
                <html>
                    <body>
                        <a href="http://%25">Invalid URL</a>
                    </body>
                </html>`,
			expected:  []string{"http://%25"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tt.inputBody, tt.inputURL)
			if (err != nil) != tt.expectErr {
				t.Errorf("getURLsFromHTML() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("getURLsFromHTML() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
