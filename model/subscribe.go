package model

type Contact struct {
	Company      string `json:"Company,omitempty"` // 收件人公司
	Name         string `json:"Name"`              // 收件人
	Tel          string `json:"Tel,omitempty"`     // 电话与手机，必填一个
	Mobile       string `json:"Mobile,omitempty"`
	PostCode     string `json:"PostCode,omitempty"` // 收件人邮编 (ShipCode 为 EMS、YZPY、YZBK 时 必填)
	ProvinceName string `json:"ProvinceName"`       // 收件省(如广东省，不要缺少“省”; 如是直辖市，请直接传北京、 上海等; 如是自治区，请直接传广西 壮族自治区等)
	CityName     string `json:"CityName"`           // 收件市(如深圳市，不要缺少 “市; 如是市辖区，请直接传北京 市、上海市等”)
	ExpAreaName  string `json:"ExpAreaName"`        // 收件区(如福田区，不要缺 少“区”或“县”）
	Address      string `json:"Address"`            // 收件人详细地址
}

type AddService struct {
	Name       string `json:"Name,omitempty"`       // 增值服务名称
	Value      string `json:"Value,omitempty"`      // 增值服务值
	CustomerID string `json:"CustomerID,omitempty"` // 客户标识(选填)
}

type Commodity struct {
	GoodsName     string `json:"GoodsName"`     // 商品名称
	GoodsCode     string `json:"GoodsCode"`     // 商品编码
	Goodsquantity string `json:"Goodsquantity"` // 件数
	GoodsPrice    string `json:"GoodsPrice"`    // 商品价格
	GoodsWeight   string `json:"GoodsWeight"`   // 商品重量kg
	GoodsDesc     string `json:"GoodsDesc"`     // 商品描述
	GoodsVol      string `json:"GoodsVol"`      // 商品体积m3
}

// SubscribeRequest 订阅 API Request
type SubscribeRequest struct {
	// Callback 用户自定义回调信息
	CallBack string `json:"Callback,omitempty"`
	// MemberID 会员标识，平台方与快递鸟统一用户标识的商家ID
	MemberID string `json:"MemberID,omitempty"`
	// WareHouseID 仓库标识(备用字段)
	WareHouseID string `json:"WareHouseID,omitempty"`
	// CustomerName 商家编码(   为 JD 时必填)
	CustomerName string `json:"CustomerName,omitempty"`
	// ShipperCode 快递公司编码
	ShipperCode string `json:"ShipperCode"`
	// LogisticCode 快递单号
	LogisticCode string `json:"LogisticCode"`
	// OrderCode 订单编号
	OrderCode string `json:"OrderCode,omitempty"`
	// MonthCode 月结编码
	MonthCode string `json:"MonthCode,omitempty"`
	// PayType 邮费支付方式: 1-现付，2-到付，3-月结，4-第三方支付(仅 SF 支持))
	PayType string `json:"PayType,omitempty"`
	// ExpType 快递类型：1-标准快件
	ExpType string `json:"ExpType,omitempty"`
	// Cost 寄件费（运费）
	Cost     string   `json:"Cost,omitempty"`
	Receiver *Contact `json:"Receiver"`
	Sender   *Contact `json:"Sender"`
	// IsNotice 是否通知快递员上门揽件：0-通知；1-不通知；不填则默认为1
	IsNotice string `json:"IsNotice,omitempty"`
	// StartDate 上门取货时间段: "yyyy-MM-dd HH:mm:ss"格式化，本文中所有时间格式相同
	StartDate string `json:"StartDate,omitempty"`
	EndDate   string `json:"EndDate,omitempty"`
	// Weight 物品总重量kg
	Weight string `json:"Weight,omitempty"`
	// Quantity 包裹数，一个包裹对应一个 运单号，如果是大于 1 个包 裹，返回则按照子母件的方 式返回母运单号和子运单号
	Quantity string `json:"Quantity,omitempty"`
	// Volume 物品总体积m3
	Volume string `json:"Volume,omitempty"`
	// Remark 备注
	Remark string `json:"Remark,omitempty"`
	// IsSendMessage 是否订阅短信：0-不需要；1-需要
	IsSendMessage string      `json:"IsSendMessage,omitempty"`
	AddService    *AddService `json:"AddService,omitempty"`
	Commodity     []Commodity `json:"Commodity,omitempty"`
}

type SubscribeResponse struct {
	EBusinessID           string `json:"EBusinessID,omitemtpy"`           // 电商用户ID
	UpdateTime            string `json:"UpdateTime,omitempty"`            // 更新时间 YYYY-MM-DD HH24:MM:SS
	Success               bool   `json:"Success"`                         // 成功与否
	Reason                string `json:"Reason,omitempty"`                // 失败原因
	EstimatedDeliveryTime string `json:"EstimatedDeliveryTime,omitempty"` // 订单预计到货时间
}

type PushData struct {
	EBusinessID string        `json:"EBusinessID"` // 电商用户ID
	PushTime    string        `json:"PushTime"`    // 推送时间
	Count       string        `json:"Count"`       // 推送轨迹的快递单号个数
	Data        []TracingData // 推送轨迹的轨迹数据集合
}

type PushResponse struct {
	EBusinessID string `json:"EBusinessID"` // 电商用户ID
	UpdateTime  string `json:"UpdateTime"`  // 更新时间 YYYY-MM-DD HH24:MM:SS
	Success     bool   `json:"Success"`     // 成功与否
}
