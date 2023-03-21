package model

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"

	"github.com/bububa/kdniao/config"
	"github.com/bububa/kdniao/util"
)

// Response general api response
type Response struct {
	RequestType string `form:"RequesType" json:"RequestType,omitempty"`
	EBusinessID string `form:"EBusinessID" json:"EBusinessID,omitempty"`
	DataSign    string `form:"DataSign" json:"DataSign,omitempty"`
	RequestData string `form:"RequestData" json:"RequestData,omitempty"`
}

func (r Response) Verify(appKey string) bool {
	data, err := url.QueryUnescape(r.RequestData)
	if err != nil {
		return false
	}
	m := Md5(util.StringsJoin(data, appKey))
	sign := base64.StdEncoding.EncodeToString([]byte(m))
	reqSign, err := url.QueryUnescape(r.DataSign)
	if err != nil {
		return false
	}
	return sign == reqSign
}

type Request struct {
	RequestData string `json:"RequestData"`
	EBusinessID string `json:"EBusinessID"`
	RequestType string `json:"RequestType"`
	DataSign    string `json:"DataSign"`
	DataType    string `json:"DataType"`
	EndPoint    string `json:"-"`
}

func NewRequest(businessId string, endPoint string) *Request {
	return &Request{
		EBusinessID: businessId,
		EndPoint:    endPoint,
	}
}

func (r *Request) SetData(requestType string, req interface{}) {
	buf, _ := json.Marshal(req)
	r.RequestData = string(buf)
	r.RequestType = requestType
}

func (r Request) Sign(appKey string) string {
	m := Md5(util.StringsJoin(r.RequestData, appKey))
	return base64.StdEncoding.EncodeToString([]byte(m))
}

func (r Request) Values(appKey string, values url.Values) {
	values.Set("EBusinessID", r.EBusinessID)
	values.Set("RequestType", r.RequestType)
	values.Set("DataSign", r.Sign(appKey))
	values.Set("RequestData", r.RequestData)
	if r.DataType != "" {
		values.Set("DataType", r.DataType)
	} else {
		values.Set("DataType", config.DEFAULT_DATATYPE)
	}
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
