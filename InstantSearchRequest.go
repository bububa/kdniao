package kdniao

type InstantSearchRequest struct {
	OrderCode    string `json:"OrderCode,omitempty"`
	ShipperCode  string `json:"ShipperCode"`
	LogisticCode string `json:"LogisticCode"`
}

func InstantSearch(c *Client, shipperCode string, logisticCode string, orderCode string) (TracingData, error) {
	req := c.NewRequest(INSTANT_SEARCH_ENDPOINT)
	apiReq := InstantSearchRequest{
		OrderCode:    orderCode,
		ShipperCode:  shipperCode,
		LogisticCode: logisticCode,
	}
	req.SetData(InstantSearchRequestType, apiReq)
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
