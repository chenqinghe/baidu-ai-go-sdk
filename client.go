package gosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const AuthUrl = "https://openapi.baidu.com/oauth/2.0/token"

// Authorizer 用于设置access_token
// 可以通过RESTFul api的方式从百度方获取
// 有效期为一个月，可以存至数据库中然后从数据库中获取
type Authorizer interface {
	AccessToken(key, secret string) (string, error)
}

type Client struct {
	key    string
	secret string

	Authorizer Authorizer

	httpclient *http.Client
}

func NewClient(ApiKey, secretKey string) *Client {
	return &Client{
		key:        ApiKey,
		secret:     secretKey,
		Authorizer: DefaultAuthorizer{},
		httpclient: &http.Client{},
	}
}

func (c *Client) SetAuther(auth Authorizer) {
	c.Authorizer = auth
}

func (c *Client) SetHttpClient(client *http.Client) {
	c.httpclient = client
}

func (c *Client) post() {

}

func (c *Client) get() {

}

type AuthResponse struct {
	AccessToken      string `json:"access_token"`  //要获取的Access Token
	ExpireIn         string `json:"expire_in"`     //Access Token的有效期(秒为单位，一般为1个月)；
	RefreshToken     string `json:"refresh_token"` //以下参数忽略，暂时不用
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	SessionSecret    string `json:"session_secret"`
	ERROR            string `json:"error"`             //错误码；关于错误码的详细信息请参考鉴权认证错误码(http://ai.baidu.com/docs#/Auth/top)
	ErrorDescription string `json:"error_description"` //错误描述信息，帮助理解和解决发生的错误。
}

type DefaultAuthorizer struct {
	tokens map[string]string
}

func (da DefaultAuthorizer) AccessToken(key, secret string) (string, error) {
	token, ok := da.tokens[key]
	if ok {
		return token, nil
	}

	query := url.Values{}
	query.Set("grant_type", "client_credentials")
	query.Set("client_id", key)
	query.Set("client_secret", secret)

	resp, err := http.Post(fmt.Sprintf("%s?%s", AuthUrl, query.Encode()), "application/x-www-form-urlencoded", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	authresponse := new(AuthResponse)
	if err := json.Unmarshal(data, authresponse); err != nil {
		return "", err
	}

	if authresponse.ERROR != "" || authresponse.AccessToken == "" {
		return "", fmt.Errorf("auth error: %s", authresponse.ErrorDescription)
	}

	da.tokens[key] = authresponse.AccessToken

	return authresponse.AccessToken, nil
}

type Params map[string]interface{}
type H = Params // H is a shorthand for Params
