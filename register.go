package park

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiRegister = "/api/tenant/register"
)

type RegisterRequest struct {
	ServiceName      string // 项目代号
	Account          string // 管理员帐号
	Password         string // 管理员密码
	TenantName       string // 商户名称
	ThirdTenantId    uint   // 本项目下的商户ID
	Url              string // 项目地址
	WxAppId          string
	WxMchId          string
	WxSubMchId 		 string
	WxKey            string
	WxPrivateKey     string
	WxCertPEM        string
	WxKeyPEM         string
	AliPayPublicKey  string
	AliPayPrivateKey string
	AliPayAppId      string
	AliPayPartnerId  string
}

type RegisterResponse struct {
	State   string
	Message string
}

func Register(req *RegisterRequest) (string, error) {
	data, _ := json.Marshal(req)
	resp, err := http.Post(req.Url+apiRegister,
		"application/json",
		strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	ret := new(RegisterResponse)
	err = json.Unmarshal(body, ret)
	if err != nil {
		return "", err
	}
	if ret.State != "success" {
		return "", errors.New(ret.Message)
	}

	return req.Url, nil
}
