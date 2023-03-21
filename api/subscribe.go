package api

import (
	"github.com/bububa/kdniao/config"
	"github.com/bububa/kdniao/core"
	"github.com/bububa/kdniao/model"
)

// Subscribe 物流跟踪
func Subscribe(c *core.Client, req *model.SearchRequest, ret *model.TracingData) error {
	r := c.NewRequest(config.SUBSCRIBE_ENDPOINT)
	r.SetData(config.SubscribeRequestType, req)
	if err := c.Request(r, ret); err != nil {
		return err
	}
	if !ret.Success {
		return model.Error{Reason: ret.Reason}
	}
	return nil
}
