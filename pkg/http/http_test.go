package http

import (
	"testing"
	"time"
)

func TestNewHTTPServerWithOptions(t *testing.T) {
	tests := []struct {
		name    string
		options []Option

		// Options to test
		expectedServerHeader  string
		expectedAppName       string
		expectedCaseSensitive bool
		expectedStrictRouting bool
		expectedReadTimeout   time.Duration
		expectedWriteTimeout  time.Duration
		expectedBodyLimit     int
	}{
		{
			name:                  "Default Config",
			options:               []Option{},
			expectedServerHeader:  "",
			expectedAppName:       "",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server Name",
			options:               []Option{WithHTTPServerHeader("test")},
			expectedServerHeader:  "test",
			expectedAppName:       "",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server and App Name",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp")},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Case Sensitive",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerCaseSensitive(true)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: true,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Strict Routing",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerStrictRouting(true)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: true,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Read Timeout",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerReadTimeout(time.Second)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   time.Second,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Write Timeout",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerWriteTimeout(time.Second)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  time.Second,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Read and Write Timeout",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerReadTimeout(time.Second), WithHTTPServerWriteTimeout(time.Second)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   time.Second,
			expectedWriteTimeout:  time.Second,
			expectedBodyLimit:     4194304,
		},
		{
			name:                  "Custom Server, App Name with Body Limit",
			options:               []Option{WithHTTPServerHeader("test"), WithHTTPServerAppName("TestApp"), WithHTTPServerBodyLimit(1024)},
			expectedServerHeader:  "test",
			expectedAppName:       "TestApp",
			expectedCaseSensitive: false,
			expectedStrictRouting: false,
			expectedReadTimeout:   0,
			expectedWriteTimeout:  0,
			expectedBodyLimit:     1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewHTTPServer(tt.options...)

			if server.App.Config().ServerHeader != tt.expectedServerHeader {
				t.Errorf("ServerHeader = %v, want %v", server.App.Config().ServerHeader, tt.expectedServerHeader)
			}

			if server.App.Config().AppName != tt.expectedAppName {
				t.Errorf("expected AppName to be %v, but got %v", tt.expectedAppName, server.App.Config().AppName)
			}

			if server.App.Config().CaseSensitive != tt.expectedCaseSensitive {
				t.Errorf("expected CaseSensitive to be %v, but got %v", tt.expectedCaseSensitive, server.App.Config().CaseSensitive)
			}

			if server.App.Config().StrictRouting != tt.expectedStrictRouting {
				t.Errorf("expected StrictRouting to be %v, but got %v", tt.expectedStrictRouting, server.App.Config().StrictRouting)
			}

			if server.App.Config().ReadTimeout != tt.expectedReadTimeout {
				t.Errorf("expected ReadTimeout to be %v, but got %v", tt.expectedReadTimeout, server.App.Config().ReadTimeout)
			}

			if server.App.Config().WriteTimeout != tt.expectedWriteTimeout {
				t.Errorf("expected WriteTimeout to be %v, but got %v", tt.expectedWriteTimeout, server.App.Config().WriteTimeout)
			}

			if server.App.Config().BodyLimit != tt.expectedBodyLimit {
				t.Errorf("expected BodyLimit to be %v, but got %v", tt.expectedBodyLimit, server.App.Config().BodyLimit)
			}
		})
	}
}
