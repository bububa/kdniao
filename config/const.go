package config

const (
	GATEWAY = "http://api.kdniao.com"
	SANDBOX = "http://sandboxapi.kdniao.com:8080/kdniaosandbox/gateway/exterfaceInvoke.json"
)

const (
	SEARCH_ENDPOINT    = "/Ebusiness/EbusinessOrderHandle.aspx"
	SUBSCRIBE_ENDPOINT = "/api/dist"
)
const (
	DEFAULT_DATATYPE = "2" // JSON
)

const (
	CONTENT_TYPE = "application/x-www-form-urlencoded;charset=utf-8"
)

const (
	EXPRESS_SF     = "SF"     // 顺丰
	EXPRESS_HTKY   = "HTKY"   // 百世快递
	EXPRESS_ZTO    = "ZTO"    // 中通
	EXPRESS_STO    = "STO"    // 申通
	EXPRESS_YTO    = "YTO"    // 圆通
	EXPRESS_YD     = "YD"     // 韵达
	EXPRESS_YZPY   = "YZPY"   // 邮政平邮
	EXPRESS_EMS    = "EMS"    // EMS
	EXPRESS_HHTT   = "HHTT"   // 天天快递
	EXPRESS_JD     = "JD"     // 京东
	EXPRESS_QFKD   = "QFKD"   // 全峰快递
	EXPRESS_GTO    = "GTO"    // 国通
	EXPRESS_UC     = "UC"     // 优速
	EXPRESS_DBL    = "DBL"    // 德邦
	EXPRESS_FAST   = "FAST"   // 快捷
	EXPRESS_AMAZON = "AMAZON" // 亚马逊
	EXPRESS_ZJS    = "ZJS"    // 宅急送
)

const (
	PayTypeSpot           = "1" // 现付
	PayTypeFreightCollect = "2" // 到付
	PayTypeMonthly        = "3" // 月结
	PayTypeThirdpary      = "4" // 第三方支付
)

const (
	ExpTypeStandard = "1"
)

const (
	StateNone            = "0" // 无轨迹
	StateTook            = "1" // 已揽收
	StateInTransit       = "2" // 在途中
	StateSignedFor       = "3" // 签收
	StateProblemShipment = "4" // 问题件
)

const (
	InstantSearchRequestType = "1002"
	SubscribeRequestType     = "1008"
	SearchMonitorRequestType = "8001"
	ExpressSearchRequestType = "8002"

	// PushTracing 轨迹查询结果
	PushTracing = "101"
	// PushPayment 货款状态
	PushPayment = "107"
)

const (
	ErrSuccess                = "100" // 成功
	ErrInsufficientParameters = "101" // 缺少必要参数
	ErrVerifyFailure          = "102" // 校验问题
	ErrInvalidFormat          = "103" // 格式问题
	ErrUserProblem            = "104" // 用户问题
	ErrOtherError             = "105" // 其他错误
	ErrReqInvalidFormat       = "401" // RequestData格式有误
	ErrNoOrder                = "402" // 缺少快递单号
	ErrSpecialChars           = "403" // 快递单号有特殊字符
	ErrOrderWrongLength       = "404" // 快递单号长度不符
	ErrExceedQueryFrequence   = "405" // 超出查询次数限制(日查询次数<=3万)
)
