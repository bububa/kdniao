package api

import (
	"github.com/bububa/kdniao/config"
	"github.com/bububa/kdniao/core"
	"github.com/bububa/kdniao/model"
)

// SearchMonitor 在途监控
func SearchMonitor(c *core.Client, req *model.SearchRequest, ret *model.TracingData) error {
	r := c.NewRequest(config.SEARCH_ENDPOINT)
	r.SetData(config.SearchMonitorRequestType, req)
	if err := c.Request(r, ret); err != nil {
		return err
	}
	if !ret.Success {
		return model.Error{Reason: ret.Reason}
	}
	return nil
}
