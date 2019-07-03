package kdniao

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type Response struct {
	RequestType string `json:"RequestType"`
	DataSign    string `json:"DataSign"`
	RequestData string `json:"RequestData"`
}

func (this Response) Verify(appKey string) bool {
	data, err := url.QueryUnescape(this.RequestData)
	if err != nil {
		return false
	}
	m := Md5(fmt.Sprintf("%s%s", data, appKey))
	sign := base64.StdEncoding.EncodeToString([]byte(m))
	reqSign, err := url.QueryUnescape(this.DataSign)
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

func (this *Request) SetData(requestType string, req interface{}) {
	buf, _ := json.Marshal(req)
	this.RequestData = string(buf)
	this.RequestType = requestType
}

func (this Request) Sign(appKey string) string {
	m := Md5(fmt.Sprintf("%s%s", this.RequestData, appKey))
	return base64.StdEncoding.EncodeToString([]byte(m))
}

func (this Request) ToValues(appKey string) url.Values {
	values := make(url.Values)
	values.Set("EBusinessID", this.EBusinessID)
	values.Set("RequestType", this.RequestType)
	values.Set("DataSign", this.Sign(appKey))
	values.Set("RequestData", this.RequestData)
	if this.DataType != "" {
		values.Set("DataType", this.DataType)
	} else {
		values.Set("DataType", DEFAULT_DATATYPE)
	}
	return values
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
