package cloudgo

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/fi-ts/cloud-go/api/client"
	"github.com/go-openapi/runtime"
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
		Timeout: c.timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	transport := httptransport.NewWithClient(c.host, c.basePath, c.schemes, client)
	transport.DefaultAuthentication = auther

	return transport
}
