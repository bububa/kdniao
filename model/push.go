package model

// PushRequest 推送API Request
type PushRequest struct {
	// EBusinessID 用户电商ID
	EBusinessID string `json:"EBusinessID,omitempty"`
	// PushTime 推送时间
	PushTime string `json:"PushTime,omitempty"`
	// Count 推送物流单号轨迹个数
	Count string `json:"Count,omitempty"`
	// Data 推送物流单号轨迹集合
	Data []TracingData `json:"Data,omitempty"`
}
