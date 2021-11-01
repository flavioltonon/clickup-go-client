package clickup

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
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

// NewClient creates a new Client with customizable http.Client and options. If a nil pointer to a http.Client is provided,
// http.DefaultClient will be used instead.
// Also, if no options are provided, the _defaultOptions will be used instead. Although, an AuthorizationMethod will always be
// required for the requests to the API to be successful.
//
// Examples:
//  client, _ := NewClient(nil) // using http.DefaultClient and _defaultOptions
//
//  auth := NewPersonalTokenAuthorization("my_token")
//  client, _ := NewClient(&http.Client{Timeout: 60 * time.Second}, SetAuthorizationMethod(auth)) // using custom http.Client and options
func NewClient(httpClient *http.Client, opts ...ClientOptionFunc) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	options := _defaultOptions

	for _, opt := range opts {
		if err := opt.apply(&options); err != nil {
			return nil, err
		}
	}

	if err := options.Validate(); err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, options: options}
	c.common.client = c
	c.Tasks = (*TasksService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	return c, nil
}

// ClientOptions are options for a Client
type ClientOptions struct {
	Authorization AuthorizationMethod
	APITargeting  *ClientAPITargetingOptions
}

// Validate validates a ClientOptions
func (o ClientOptions) Validate() error {
	return ozzo.ValidateStruct(&o,
		ozzo.Field(&o.Authorization, ozzo.Required),
		ozzo.Field(&o.APITargeting, ozzo.Required),
	)
}

// ClientAPITargetingOptions are ClientOptions about ClickUp's API
type ClientAPITargetingOptions struct {
	Scheme string
	Host   string
}

// Validate validates a ClientAPITargetingOptions
func (o ClientAPITargetingOptions) Validate() error {
	return ozzo.ValidateStruct(&o,
		ozzo.Field(&o.Scheme, ozzo.Required),
		ozzo.Field(&o.Host, ozzo.Required),
	)
}

// _defaultOptions are the default options for Clients
var _defaultOptions = ClientOptions{
	APITargeting: &ClientAPITargetingOptions{
		Scheme: "https",
		Host:   "api.clickup.com",
	},
}

// ClientOptionFunc is an option for a Client
type ClientOptionFunc func(c *ClientOptions) error

// apply applies the ClientOptionFunc to a Client
func (fn ClientOptionFunc) apply(c *ClientOptions) error { return fn(c) }

// SetAuthorizationMethod sets the client options with an authorization method.
func SetAuthorizationMethod(authorization AuthorizationMethod) ClientOptionFunc {
	return func(o *ClientOptions) error {
		o.Authorization = authorization
		return nil
	}
}

// SetPersonalTokenAuthorization sets the client options with a personal token authorization method.
func SetPersonalTokenAuthorization(token string) ClientOptionFunc {
	return func(o *ClientOptions) error {
		o.Authorization = NewPersonalTokenAuthorization(token)
		return nil
	}
}

// SetAPIURL sets a new URL to the client options to target ClickUp's API.
func SetAPIURL(apiURL string) ClientOptionFunc {
	return func(o *ClientOptions) error {
		u, err := url.Parse(apiURL)
		if err != nil {
			return err
		}

		o.APITargeting = &ClientAPITargetingOptions{
			Scheme: u.Scheme,
			Host:   u.Host,
		}
		return nil
	}
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
	request.SetAuthorization(c.options.Authorization)
	return request, nil
}

// SetHeader sets a key/value pair in the request headers
func (r *Request) SetHeader(key, value string) { r.Header.Set(key, value) }

// SetContentType sets a Content-Type header with a given value in the request
func (r *Request) SetContentType(contentType string) { r.SetHeader(header_contentType, contentType) }

// SetAuthorization sets an Authorization header in the request for a given AuthorizationMethod
func (r *Request) SetAuthorization(method AuthorizationMethod) { method.ApplyAuthorization(r) }

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
