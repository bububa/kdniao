package api

import (
	"github.com/bububa/kdniao/config"
	"github.com/bububa/kdniao/core"
	"github.com/bububa/kdniao/model"
)

// ExpressSearch 快递查询
func ExpressSearch(c *core.Client, req *model.SearchRequest, ret *model.TracingData) error {
	r := c.NewRequest(config.SEARCH_ENDPOINT)
	r.SetData(config.ExpressSearchRequestType, req)
	if err := c.Request(r, ret); err != nil {
		return err
	}
	if !ret.Success {
		return model.Error{Reason: ret.Reason}
	}
	return nil
}
