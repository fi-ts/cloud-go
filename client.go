package cloud

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/metal-pod/cloud-go/api/client"
	"github.com/metal-pod/security"

	httptransport "github.com/go-openapi/runtime/client"
)

// NewClient creates a new client for accessing the cloud-api
func NewClient(rawurl, token, hmac string) (*client.Cloud, error) {
	if (token == "") == (hmac == "") {
		return nil, errors.New("either token or hmac is required")
	}

	parsedurl, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if parsedurl.Host == "" {
		return nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", rawurl)
	}

	var hmacAuth *security.HMACAuth
	if hmac != "" {
		auth := security.NewHMACAuth("Metal-View", []byte(hmac))
		hmacAuth = &auth
	}

	auther := runtime.ClientAuthInfoWriterFunc(func(rq runtime.ClientRequest, rg strfmt.Registry) error {
		if hmacAuth != nil {
			hmacAuth.AddAuthToClientRequest(rq, time.Now())
		} else if token != "" {
			security.AddUserTokenToClientRequest(rq, token)
		}
		return nil
	})

	transport := httptransport.New(parsedurl.Host, parsedurl.Path, []string{parsedurl.Scheme})
	transport.DefaultAuthentication = auther

	c := client.New(transport, nil)

	return c, nil
}
