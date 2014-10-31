package semantics3

import (
	"encoding/json"
	"errors"
	"github.com/mrjones/oauth"
	"net/http"
)

var (
	Host = "https://api.semantics3.com/v1/"
)

var ErrOAuthRequired = errors.New("OAuth is required")

type OAuthConsumer interface {
	Get(string, map[string]string, *oauth.AccessToken) (*http.Response, error)
}

type Client struct {
	endpoint  string
	dataquery string
	oauth     OAuthConsumer
}

func NewClient(api_key, api_secret, endpoint string) *Client {
	oa := oauth.NewConsumer(api_key, api_secret, oauth.ServiceProvider{})
	client := &Client{endpoint: endpoint, oauth: oa}
	return client
}

func (c *Client) AddParams(params map[string]interface{}) {
	pjson, _ := json.Marshal(params)
	c.dataquery = string(pjson)
}

func (c *Client) getBaseUrl() string {
	return Host + c.endpoint
}

func (c *Client) getParams() string {
	return c.dataquery
}

func (c *Client) Get() (*http.Response, error) {
	if c.oauth == nil {
		return nil, ErrOAuthRequired
	}
	baseurl := c.getBaseUrl()
	params := c.getParams()
	response, err := c.oauth.Get(baseurl, map[string]string{"q": params}, &oauth.AccessToken{})
	return response, err
}
