package core

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bububa/kdniao/config"
	"github.com/bububa/kdniao/model"
	"github.com/bububa/kdniao/util"
)

type Client struct {
	EBusinessID string
	AppKey      string
	IsSandbox   bool
	clt         *http.Client
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
		clt:         c,
	}
}

func (c *Client) Sandbox() {
	c.IsSandbox = true
}

func (c *Client) Production() {
	c.IsSandbox = false
}

func (c *Client) gateway() string {
	if c.IsSandbox {
		return config.SANDBOX
	}
	return config.GATEWAY
}

func (c *Client) NewRequest(endPoint string) *model.Request {
	return model.NewRequest(c.EBusinessID, endPoint)
}

func (c *Client) Request(req *model.Request, rep interface{}) error {
	var uri string
	if c.IsSandbox {
		uri = c.gateway()
	} else {
		uri = util.StringsJoin(c.gateway(), req.EndPoint)
	}
	values := util.GetUrlValues()
	req.Values(c.AppKey, values)
	resp, err := c.clt.Post(uri, config.CONTENT_TYPE, strings.NewReader(values.Encode()))
	util.PutUrlValues(values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(&rep)
}
