package kdniao

type TracingData struct {
	EBusinessID           string        `json:"EBusinessID,omitemtpy"`           // 电商用户ID
	OrderCode             string        `json:"OrderCode,omitempty"`             // 订单编号
	ShipperCode           string        `json:"ShipperCode"`                     // 快递公司编码
	LogisticCode          string        `json:"LogisticCode"`                    // 物流运单号
	Success               bool          `json:"Success"`                         // 成功与否
	Reason                string        `json:"Reason,omitempty"`                // 失败原因
	State                 string        `json:"State"`                           // 物流状态: 0-无轨迹 1-已揽收 2-在途中 3-签收 4-问题件
	CallBack              string        `json:"CallBack,omitempty"`              // 订阅接口的Bk值
	Traces                []TracingItem `json:"Traces"`                          // 物流轨迹详情
	EstimatedDeliveryTime string        `json:"EstimatedDeliveryTime,omitempty"` // 预计到达时间yyyy-mm-dd
	PickerInfo            struct {
		PersonName     string `json:"PersonName,omitempty"`     // 快递员姓名
		PersonTel      string `json:"PersonTel,omitempty"`      // 快递员电话
		PersonCode     string `json:"PersonCode,omitempty"`     // 快递员工号
		StationName    string `json:"StationName,omitempty"`    // 网点名称
		StationAddress string `json:"StationAddress,omitempty"` // 网点地址
		StationTel     string `json:"StationTel,omitempty"`     // 网点电话
	} `json:"PickerInfo,omitempty"` // 收件员信息
	SenderInfo struct {
		PersonName     string `json:"PersonName,omitempty"`     // 派件员姓名
		PersonTel      string `json:"PersonTel,omitempty"`      // 派件员快递员电话
		PersonCode     string `json:"PersonCode,omitempty"`     // 派件员快递员工号
		StationName    string `json:"StationName,omitempty"`    // 派件员网点名称
		StationAddress string `json:"StationAddress,omitempty"` // 派件员网点地址
		StationTel     string `json:"StationTel,omitempty"`     // 派件员网点电话
	} `json:"SenderInfo,omitempty"` // 派件员信息
}

type TracingItem struct {
	AcceptTime    string `json:"AcceptTime"`       // 时间
	AcceptStation string `json:"AcceptStation"`    // 描述
	Remark        string `json:"Remark,omitempty"` // 备注
}
