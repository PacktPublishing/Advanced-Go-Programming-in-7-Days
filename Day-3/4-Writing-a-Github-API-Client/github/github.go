package github

import (
	"net/http"
	"net/url"
	"strings"
	"io"
	"bytes"
	"encoding/json"
	"context"
	"golang.org/x/oauth2"
	"io/ioutil"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
)

const (
	defaultBaseURL = "https://api.github.com/"
	acceptVersionHeader = "application/vnd.github.v3+json" // https://developer.github.com/v3/#current-version
)

// A Client manages communication with the Github API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to the public Github API
	baseURL *url.URL

	// Services used for handling the Github API.
	Repositories *RepositoriesService
}

// NewClient returns a new GitLab API client. You must provide a valid token.
func NewClient(ctx context.Context, token string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := newClient(tc)
	return client
}

func newClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, UserAgent: userAgent}
	c.SetBaseURL(defaultBaseURL)

	// Create all the public services.
	c.Repositories = &RepositoriesService{client:c}

	return c
}

// GetBaseURL returns a copy of the baseURL.
func (c *Client) GetBaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewRequest creates an API request using a relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	buf, err := c.encodeRequestBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// Set version header
	req.Header.Set("Accept", acceptVersionHeader)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := ctxhttp.Do(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		rerr := resp.Body.Close()
		if rerr != nil {
			err = rerr
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}


	err = c.decodeResponseBody(resp.Body, v)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CheckResponse checks the API response for errors,
// and returns them if present.
func CheckResponse(r *http.Response) error {
	if r.StatusCode < 300 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(r.Response.Request.URL.Opaque)
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, path,
		r.Response.StatusCode, r.Message)
}

func (c *Client) decodeResponseBody(body io.ReadCloser, v interface{}) (error) {
	if v != nil {
		var err error
		err = json.NewDecoder(body).Decode(v)
		return err
	}
	return nil
}

func (c *Client) encodeRequestBody(body interface{}) (io.ReadWriter, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		} else {
			return buf, nil
		}
	} else {
		return nil, nil
	}
}