package semantics3

import (
    "encoding/json"
    "net/http"
    "github.com/mrjones/oauth"
)

var (
    Host = "https://api.semantics3.com/v1/"
)

type Client struct {
    api_key     string
    api_secret  string
    endpoint    string
    dataquery   string
}

func NewClient(api_key, api_secret, endpoint string) *Client {
    return &Client{api_key: api_key, api_secret: api_secret, endpoint: endpoint}
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
    baseurl := c.getBaseUrl()
    params := c.getParams()
    oa := oauth.NewConsumer(c.api_key, c.api_secret, oauth.ServiceProvider{})
    response, err := oa.Get(baseurl, map[string]string{"q": params}, &oauth.AccessToken{})
    return response, err
}

