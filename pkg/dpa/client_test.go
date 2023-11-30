package dpa

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests cases used in multiple tests
var Tests = []struct {
	name    string
	header  http.ConnState
	path    string
	payload interface{}
	wantErr bool
}{
	{
		name:    "Invalid - GatewayTimeout",
		header:  http.StatusGatewayTimeout,
		payload: map[string]string{"example": "valid"},
		wantErr: true,
	},
	{
		name:    "Valid - StatusOK",
		header:  http.StatusOK,
		payload: map[string]string{"example": "valid"},
		wantErr: false,
	},
	{
		name:   "Invalid - Payload Marshal Failure",
		header: http.StatusOK,
		payload: map[string]interface{}{
			"example": make(chan int),
		},
		wantErr: true,
	},
	{
		name:    "Invalid - Invalid Character in Path",
		path:    "/foo\bar",
		header:  0,
		wantErr: true,
	},
}

func TestGet(t *testing.T) {
	// Get does not support payload so it needs standard test cases
	var tests = []struct {
		name     string
		header   http.ConnState
		payload  interface{}
		path     string
		response string
		wantErr  bool
	}{
		{
			name:     "Invalid - GatewayTimeout",
			header:   http.StatusGatewayTimeout,
			response: "error",
			wantErr:  true,
		},
		{
			name:     "Valid - StatusOK",
			header:   http.StatusOK,
			response: "ok",
			wantErr:  false,
		},
		{
			name:    "Invalid - Invalid Character in Path",
			path:    "/foo\bar",
			header:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if tt.header != 0 {
					w.WriteHeader(int(tt.header))
				}
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			err := client.Get(context.Background(), tt.path, nil, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestPost(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			err := client.Post(context.Background(), tt.path, tt.payload, nil, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestPut(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			err := client.Put(context.Background(), tt.path, tt.payload, nil, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestPatch(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			err := client.Patch(context.Background(), tt.path, tt.payload, nil, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Patch() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("Patch() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			err := client.Delete(context.Background(), tt.path, tt.payload, nil, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	var tests = []struct {
		name    string
		header  http.ConnState
		path    string
		payload interface{}
		wantErr bool
	}{
		{
			name:    "Valid - StatusOK",
			header:  http.StatusOK,
			payload: map[string]string{"example": "valid"},
			wantErr: false,
		},
		{
			name:   "Invalid - Payload Marshal Failure",
			header: http.StatusOK,
			payload: map[string]interface{}{
				"example": make(chan int),
			},
			wantErr: true,
		},
		{
			name:    "Invalid - Invalid Character in Path",
			path:    "/foo\bar",
			header:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			_, err := client.newRequest(context.Background(), "POST", tt.path, tt.payload)
			if tt.wantErr {
				if err == nil {
					t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestNewRequest_Verbose(t *testing.T) {
	// Mock Response with expected error
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"Example":"Body"`))
	}))

	defer ts.Close()
	client := NewClient(
		&http.Client{},
		Options{
			ApiURL:  ts.URL,
			Verbose: true,
		},
	)

	req, _ := client.newRequest(context.Background(), "POST", "", map[string]string{"example": "valid"})
	if req.Body == nil {
		t.Errorf("got no req body in verbose log when req body was provided")
	}
}

func TestDoRequest(t *testing.T) {
	var tests = []struct {
		name     string
		header   http.ConnState
		response interface{}
		payload  interface{}
		wantErr  bool
	}{
		{
			name:     "Valid - StatusOK",
			header:   http.StatusOK,
			response: struct{ Example string }{},
			payload:  map[string]string{"example": "valid"},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response with expected error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if tt.header != 0 {
					w.WriteHeader(int(tt.header))
				}
				w.Write([]byte(`"Example":"Body"`))
			}))

			defer ts.Close()
			client := NewClient(
				&http.Client{},
				Options{
					ApiURL:  ts.URL,
					Verbose: false,
				},
			)

			var reqBody io.Reader
			req, _ := http.NewRequest(http.MethodGet, ts.URL, reqBody)

			err := client.doRequest(req, tt.response, nil)
			if tt.wantErr {
				if err == nil {
					t.Errorf("DoRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("DoRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestDo(t *testing.T) {
	// Mock Response with expected error
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"Example":"Body"`))
	}))

	defer ts.Close()
	client := NewClient(
		&http.Client{},
		Options{
			ApiURL:  ts.URL,
			Verbose: true,
		},
	)

	var reqBody io.Reader
	req, _ := http.NewRequest(http.MethodGet, ts.URL, reqBody)

	resp, _, _ := client.do(req)
	if resp.Body == nil {
		t.Errorf("got no body in verbose log when body was returned")
	}
}
