package clickup

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is a ClickUp API client
type Client struct {
	Tasks *TasksService
	Teams *TeamsService

	client  *http.Client
	common  service
	options ClientOptions
}

type service struct {
	client *Client
}

// NewClient creates a new Client with a customizable http.Client. If a nil pointer to a http.Client is provided,
// http.DefaultClient will be used instead.
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, options: _defaultOptions}
	c.common.client = c
	c.Tasks = (*TasksService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	return c, nil
}

// ClientOptions are options for a Client
type ClientOptions struct {
	authorization AuthorizationMethod
}

var _defaultOptions = ClientOptions{}

// SetAuthorizationMethod sets the client with an authorization method.
func (c *Client) SetAuthorizationMethod(authorization AuthorizationMethod) {
	c.options.authorization = authorization
}

// Request is a request sent to ClickUp's API
type Request struct {
	*http.Request
}

// NewRequest creates a new client Request
func (c *Client) NewRequest(method string, url *url.URL, body io.Reader) (*Request, error) {
	rawRequest, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	request := &Request{rawRequest}
	request.SetContentType("application/json")
	request.SetAuthorization(c.options.authorization)
	return request, nil
}

// SetHeader sets a key/value pair in the request headers
func (r *Request) SetHeader(key, value string) { r.Header.Set(key, value) }

// SetContentType sets a Content-Type header with a given value in the request
func (r *Request) SetContentType(contentType string) {
	r.SetHeader(_contentType.String(), contentType)
}

// SetAuthorization sets an Authorization header in the request for a given AuthorizationMethod
func (r *Request) SetAuthorization(method AuthorizationMethod) {
	r.SetHeader(_authorization.String(), method.String())
}

// Do sends a given Request and returns a Response. The response body closure is already handled and does not need to be closed.
func (c *Client) Do(request *Request) (*Response, error) {
	r, err := c.client.Do(request.Request)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	response, err := NewResponse(r)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Response is a response from ClickUp's API
type Response struct {
	RawResponse *http.Response

	Body []byte
}

// NewResponse creates a new Response from a http.Response pointer, reading its body in the process.
func NewResponse(r *http.Response) (*Response, error) {
	response := &Response{
		RawResponse: r,
	}

	if response.isSuccess() {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		response.Body = body
		return response, nil
	}

	wrapper := new(struct {
		Message string `json:"err"`
		Code    string `json:"ECODE"`
	})

	if err := response.Decode(&wrapper); err != nil {
		return nil, err
	}

	return nil, NewError(response.StatusCode(), wrapper.Message)
}

// StatusCode returns the status code of a Response
func (r *Response) StatusCode() int {
	return r.RawResponse.StatusCode
}

func (r *Response) isSuccess() bool {
	return r.StatusCode() >= 200 && r.StatusCode() < 300
}

// Decode JSON unmarshals the response body into a given value pointer
func (r *Response) Decode(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// Get sents a GET request to a given URL and returns its Response
func (c *Client) Get(url *url.URL) (*Response, error) {
	request, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(request)
}

// Get sents a POST request to a given URL with a given body and returns its Response
func (c *Client) Post(url *url.URL, body interface{}) (*Response, error) {
	var buffer bytes.Buffer

	if err := json.NewEncoder(&buffer).Encode(body); err != nil {
		return nil, err
	}

	request, err := c.NewRequest(http.MethodPost, url, &buffer)
	if err != nil {
		return nil, err
	}

	return c.Do(request)
}

// Get sents a DELETE request to a given URL and returns an error in case it fails
func (c *Client) Delete(url *url.URL) error {
	request, err := c.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if _, err := c.Do(request); err != nil {
		return err
	}

	return nil
}
