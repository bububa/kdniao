package model

import (
	"sync"
)

var tracingDataPool = sync.Pool{
	New: func() any {
		return new(TracingData)
	},
}

func GetTracingDataInPool() *TracingData {
	return tracingDataPool.Get().(*TracingData)
}

func PutTracingDataToPool(v *TracingData) {
	v.Reset()
	tracingDataPool.Put(v)
}

// SearchRequest 即时查询API Request
type SearchRequest struct {
	// OrderCode 订单编号
	OrderCode string `json:"OrderCode,omitempty"`
	// ShipperCode 快递公司编码
	ShipperCode string `json:"ShipperCode,omitempty"`
	// LogisticCode 物流单号
	LogisticCode string `json:"LogisticCode,omitempty"`
	// CustomerName ShipperCode为SF时必填，对应寄件人/收件人手机号后四位；ShipperCode为其他快递时，可不填或保留字段，不可传值
	CustomerName string `json:"CustomerName,omitempty"`
	// Callback 户自定义回传字段
	Callback string `json:"Callback,omitempty"`
}

// TracingData 物流轨迹
type TracingData struct {
	// EBusinessID 电商用户ID
	EBusinessID string `json:"EBusinessID,omitemtpy"`
	// OrderCode 订单编号
	OrderCode string `json:"OrderCode,omitempty"`
	// ShipperCode 快递公司编码
	ShipperCode string `json:"ShipperCode"`
	// LogisticCode 物流运单号
	LogisticCode string `json:"LogisticCode"`
	// Success 成功与否
	Success bool `json:"Success"`
	// Reason 失败原因
	Reason string `json:"Reason,omitempty"`
	// State 物流状态: 0-无轨迹 1-已揽收 2-在途中 3-签收 4-问题件
	State string `json:"State"`
	// StateEx 增值物流状态： 1-已揽收， 2-在途中， 201-到达派件城市， 202-派件中， 211-已放入快递柜或驿站， 3-已签收， 311-已取出快递柜或驿站， 4-问题件， 401-发货无信息， 402-超时未签收， 403-超时未更新， 404-拒收（退件）， 412-快递柜或驿站超时未取
	StateEx string `json:"StateEx,omitempty"`
	// Location 增值所在城市
	Location string `json:"location,omitempty"`
	// Traces 物流轨迹详情
	Traces []TracingItem `json:"Traces"`
	// CallBack 订阅接口的Bk值
	CallBack string `json:"CallBack,omitempty"`
	// EstimatedDeliveryTime 预计到达时间yyyy-mm-dd
	EstimatedDeliveryTime string `json:"EstimatedDeliveryTime,omitempty"`
	// PickerInfo 收件员信息
	PickerInfo *DeliverContact `json:"PickerInfo,omitempty"`
	// SenderInfo 派件员信息
	SenderInfo *DeliverContact `json:"SenderInfo,omitempty"`
	// OrderState 订单货款状态：1-待出款；2-已出款；3-已收款
	OrderState string `json:"OrderState,omitempty"`
	// AccountName 返款银行卡开户人（例：**伟、*佳）
	AccountName string `json:"AccountName,omitempty"`
	// AccountTel 返款银行卡手机末四位
	AccountTel string `json:"AccountTel,omitempty"`
	// AccountNum 返款银行卡末四位
	AccountNum string `json:"AccountNum,omitempty"`
}

func (t *TracingData) Reset() {
	t.EBusinessID = ""
	t.OrderCode = ""
	t.ShipperCode = ""
	t.LogisticCode = ""
	t.Success = false
	t.Reason = ""
	t.State = ""
	t.StateEx = ""
	t.Location = ""
	t.Traces = nil
	t.CallBack = ""
	t.EstimatedDeliveryTime = ""
	t.PickerInfo = nil
	t.SenderInfo = nil
	t.OrderState = ""
	t.AccountName = ""
	t.AccountTel = ""
	t.AccountNum = ""
}

// TracingItem 物流轨迹详情
type TracingItem struct {
	// AcceptTime 时间
	AcceptTime string `json:"AcceptTime"`
	// AcceptStation 描述
	AcceptStation string `json:"AcceptStation"`
	// Remark 备注
	Remark string `json:"Remark,omitempty"`
	// Location 当前城市
	Location string `json:"Location,omitempty"`
	// Action 当前状态
	Action string `json:"Action,omitempty"`
}

// DeliverContact 收件员信息
type DeliverContact struct {
	// PersonName 快递员姓名
	PersonName string `json:"PersonName,omitempty"`
	// PersonTel 快递员电话
	PersonTel string `json:"PersonTel,omitempty"`
	// PersonCode 快递员工号
	PersonCode string `json:"PersonCode,omitempty"`
	// StationName 网点名称
	StationName string `json:"StationName,omitempty"`
	// StationAddress 网点地址
	StationAddress string `json:"StationAddress,omitempty"`
	// StationTel 网点电话
	StationTel string `json:"StationTel,omitempty"`
}
