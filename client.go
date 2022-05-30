package cloudgo

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fi-ts/cloud-go/api/client"
	"github.com/go-openapi/runtime"
	openclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/metal-stack/security"

	httptransport "github.com/go-openapi/runtime/client"
)

const (
	defaultHMACAuthType = "Metal-View-All"
)

type clientSpec struct {
	host     string
	basePath string
	schemes  []string

	token string
	hmac  string

	hmacAuthType string
	timeout      time.Duration
}

type option func(c *clientSpec)

// AuthType sets the authType for HMAC-Auth
func AuthType(authType string) option {
	return func(c *clientSpec) {
		c.hmacAuthType = authType
	}
}

//Timeout sets the timeout for a new client
func Timeout(timeout time.Duration) option {
	// overriding the default timeout as otherwise all request need to be called with
	// WithTimeout or WithContext
	// still open issue: https://github.com/go-swagger/go-swagger/issues/2583
	openclient.DefaultTimeout = timeout

	return func(c *clientSpec) {
		c.timeout = timeout
	}
}

// NewClient creates a new client for accessing the cloud-api
func NewClient(rawurl, token, hmac string, options ...option) (*client.CloudAPI, error) {
	parsedurl, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if parsedurl.Host == "" {
		return nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", rawurl)
	}

	spec := &clientSpec{
		host:         parsedurl.Host,
		basePath:     parsedurl.Path,
		schemes:      []string{parsedurl.Scheme},
		token:        token,
		hmac:         hmac,
		hmacAuthType: defaultHMACAuthType,
		timeout:      30 * time.Second,
	}

	for _, opt := range options {
		opt(spec)
	}

	transport := newClientTransport(spec)

	c := client.New(transport, nil)

	return c, nil
}

func newClientTransport(c *clientSpec) *httptransport.Runtime {
	var hmacAuth *security.HMACAuth
	if c.hmac != "" {
		auth := security.NewHMACAuth(c.hmacAuthType, []byte(c.hmac))
		hmacAuth = &auth
	}

	auther := runtime.ClientAuthInfoWriterFunc(func(rq runtime.ClientRequest, rg strfmt.Registry) error {
		if hmacAuth != nil {
			hmacAuth.AddAuthToClientRequest(rq, time.Now())
		} else if c.token != "" {
			security.AddUserTokenToClientRequest(rq, c.token)
		}
		return nil
	})

	client := &http.Client{
		Timeout:   c.timeout,
		Transport: http.DefaultTransport,
	}

	transport := httptransport.NewWithClient(c.host, c.basePath, c.schemes, client)
	transport.DefaultAuthentication = auther

	return transport
}
