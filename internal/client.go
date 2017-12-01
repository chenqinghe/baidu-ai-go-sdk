package internal

import (
	"errors"

	"github.com/imroc/req"
)

const VOICE_AUTH_URL string = "https://openapi.baidu.com/oauth/2.0/token"

type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Authorizer   Authorizer
}

//授权成功响应信息
type AuthResponseSuccess struct {
	AccessToken   string `json:"access_token"`  //要获取的Access Token
	ExpireIn      string `json:"expire_in"`     //Access Token的有效期(秒为单位，一般为1个月)；
	RefreshToken  string `json:"refresh_token"` //以下参数忽略，暂时不用
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

//授权失败响应信息
type AuthResponseFailed struct {
	ERROR            string `json:"error"`             //错误码；关于错误码的详细信息请参考鉴权认证错误码(http://ai.baidu.com/docs#/Auth/top)
	ErrorDescription string `json:"error_description"` //错误描述信息，帮助理解和解决发生的错误。
}

//Authorizer 用于设置access_token
//可以通过RESTFul api的方式从百度方获取
//有效期为一个月，可以存至数据库中然后从数据库中获取
type Authorizer interface {
	Authorize(client *Client) error
}

type DefaultAuthorizer struct{}

type RestApiAuthorizer DefaultAuthorizer

func (da DefaultAuthorizer) Authorize(client *Client) error {
	resp, err := req.Post(VOICE_AUTH_URL, req.Param{
		"grant_type":    "client_credentials",
		"client_id":     client.ClientID,
		"client_secret": client.ClientSecret,
	})
	if err != nil {
		return err
	}
	var rsSuccess AuthResponseSuccess
	var rsFail AuthResponseFailed
	if err := resp.ToJSON(&rsSuccess); err != nil || rsSuccess.AccessToken == "" { //json解析失败
		if err := resp.ToJSON(&rsFail); err != nil || rsFail.ERROR == "" { //json解析失败
			return errors.New("授权信息解析失败:" + err.Error())
		}
		return errors.New("授权失败:" + rsFail.ErrorDescription)
	}
	client.AccessToken = rsSuccess.AccessToken
	return nil
}

func (client *Client) Auth() error {
	if client.AccessToken != "" {
		return nil
	}
	return client.Authorizer.Authorize(client)
}

func (client *Client) SetAuther(auth Authorizer) {
	client.Authorizer = auth
}

func NewClient(ApiKey, secretKey string) *Client {
	return &Client{
		ClientID:     ApiKey,
		ClientSecret: secretKey,
		Authorizer:   DefaultAuthorizer{},
	}
}
