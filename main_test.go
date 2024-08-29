package main

import (
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		expected   string
		expectExit bool
	}{
		{
			name:       "no website provided",
			args:       []string{},
			expected:   "no website provided\n",
			expectExit: true,
		},
		{
			name:       "too many arguments provided",
			args:       []string{"https://example.com", "https://another.com"},
			expected:   "too many arguments provided\n",
			expectExit: true,
		},
		{
			name:       "valid website provided",
			args:       []string{"https://example.com"},
			expected:   "starting crawl of: https://example.com\n",
			expectExit: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origArgs := os.Args
			defer func() { os.Args = origArgs }()

			os.Args = append([]string{"cmd"}, tt.args...)

			exitCode := 0
			defer func() {
				if r := recover(); r != nil {
					if code, ok := r.(int); ok {
						exitCode = code
					} else {
						t.Fatalf("unexpected panic: %v", r)
					}
				}
				if exitCode == 0 && tt.expectExit {
					t.Errorf("expected exit with non-zero status, got exit code %d", exitCode)
				} else if exitCode != 0 && !tt.expectExit {
					t.Errorf("expected no exit, got exit code %d", exitCode)
				}
			}()
			parseArgs()
		})
	}
}
