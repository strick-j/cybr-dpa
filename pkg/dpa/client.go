package dpa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type Options struct {
	ApiURL  string
	Verbose bool
}

type Client struct {
	httpClient *http.Client
	options    *Options
}

var (
	ErrUserAccessDenied = errors.New("you do not have access to the requested resource")
	ErrNotFound         = errors.New("the requested resource not found")
	ErrTooManyRequests  = errors.New("you have exceeded throttle")
)

func NewClient(httpClient *http.Client, options Options) *Client {
	return &Client{
		httpClient: httpClient,
		options:    &options,
	}
}

// Basic Interface Definitions
type HTTPClient interface {
	Get(ctx context.Context, path string, v interface{}, e interface{}) error
	Post(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error
	Put(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error
	Patch(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error
	Delete(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error
}

// //////////// COMMON METHODS - GET, POST, PUT, PATCH,  DELETE ///////////////////////////////////////////////
func (c *Client) Get(ctx context.Context, path string, v interface{}, e interface{}) error {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	if err := c.doRequest(req, v, e); err != nil {
		return err
	}

	return nil
}

func (c *Client) Post(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error {
	req, err := c.newRequest(ctx, http.MethodPost, path, payload)
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	if err := c.doRequest(req, v, e); err != nil {
		return err
	}

	return nil
}

func (c *Client) Put(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error {
	req, err := c.newRequest(ctx, http.MethodPut, path, payload)
	if err != nil {
		return fmt.Errorf("failed to create PUT request: %w", err)
	}

	if err := c.doRequest(req, v, e); err != nil {
		return err
	}

	return nil
}

func (c *Client) Patch(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error {
	req, err := c.newRequest(ctx, http.MethodPatch, path, payload)
	if err != nil {
		return fmt.Errorf("failed to create PATCH request: %w", err)
	}

	if err := c.doRequest(req, v, e); err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, path string, payload interface{}, v interface{}, e interface{}) error {
	req, err := c.newRequest(ctx, http.MethodDelete, path, payload)
	if err != nil {
		return fmt.Errorf("failed to create DELETE request: %w", err)
	}

	if err := c.doRequest(req, v, e); err != nil {
		return err
	}

	return nil
}

////////////// REQUEST PROCESSING - newRequest, doRequest, do ///////////////////////////////////////////////

func (c *Client) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
	var reqBody io.Reader
	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.options.ApiURL, path), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	if c.options.Verbose {
		body, _ := httputil.DumpRequest(req, true)
		log.Println(string(body))
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) doRequest(r *http.Request, v interface{}, e interface{}) error {
	resp, ok, err := c.do(r)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	// If not OK try to decode response body into error interface
	if !ok {
		var buf bytes.Buffer
		dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
		if err := dec.Decode(e); err != nil {
			return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
		}
		return nil
		// If OK try to decode response body into supplied response interface
	} else if ok {
		switch v := v.(type) {
		case nil:
			return nil
		case *string:
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("could not read response body: %w [%s:%s]", err, r.Method, r.URL.String())
			}
			*v = string(b)
			return nil
		default:
			var buf bytes.Buffer
			dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
			if err := dec.Decode(v); err != nil {
				return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
			}
		}
		return nil
	}
	return nil
}

func (c *Client) do(r *http.Request) (*http.Response, bool, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, false, fmt.Errorf("failed to make request [%s:%s]: %w", r.Method, r.URL.String(), err)
	}

	if c.options.Verbose {
		body, _ := httputil.DumpResponse(resp, true)
		log.Println(string(body))
	}

	switch resp.StatusCode {
	case http.StatusOK,
		http.StatusCreated,
		http.StatusMultiStatus,
		http.StatusNoContent:
		return resp, true, nil
	case http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusInternalServerError:
		return resp, false, nil
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		return nil, false, ErrTooManyRequests
	}

	return nil, false, fmt.Errorf("failed to do request, %d status code received", resp.StatusCode)
}
