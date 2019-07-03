package kdniao

func Subscribe(c *Client, apiReq SubscribeRequest) (TracingData, error) {
	req := c.NewRequest(SUBSCRIBE_ENDPOINT)
	req.SetData(SubscribeRequestType, apiReq)
	var resp TracingData
	err := c.Request(req, &resp)
	if err != nil {
		return resp, err
	}
	if !resp.Success {
		return resp, Error{Reason: resp.Reason}
	}
	return resp, err
}
