package kdniao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	EBusinessID string
	AppKey      string
	IsSandbox   bool
	*http.Client
}

func NewClient(bussinessId string, appKey string, httpClient ...*http.Client) *Client {
	var c *http.Client
	if len(httpClient) > 0 {
		c = httpClient[0]
	} else {
		c = http.DefaultClient
	}
	return &Client{
		EBusinessID: bussinessId,
		AppKey:      appKey,
		Client:      c,
	}
}

func (this *Client) Sandbox() {
	this.IsSandbox = true
}

func (this *Client) Production() {
	this.IsSandbox = false
}

func (this *Client) gateway() string {
	if this.IsSandbox {
		return SANDBOX
	}
	return GATEWAY
}

func (this *Client) NewRequest(endPoint string) *Request {
	return NewRequest(this.EBusinessID, endPoint)
}

func (this *Client) Request(req *Request, rep interface{}) error {
	var uri string
	if this.IsSandbox {
		uri = this.gateway()
	} else {
		uri = fmt.Sprintf("%s%s", this.gateway(), req.EndPoint)
	}
	values := req.ToValues(this.AppKey)
	r, err := this.Post(uri, CONTENT_TYPE, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&rep)
}
