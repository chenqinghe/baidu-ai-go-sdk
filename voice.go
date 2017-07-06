// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package dueros

import (
	"errors"

	"io/ioutil"

	"github.com/imroc/req"
)

const TSN_URL string = "http://tsn.baidu.com/text2audio"
const VOP_URL string = "http://vop.baidu.com/server_api"
const VOICE_AUTH_URL string = "https://openapi.baidu.com/oauth/2.0/token"

type TSNRequest struct {
}
type TSNResponse struct {
}

type VOPRequest struct {
}
type VOPResponse struct {
}

type AuthResponse struct {
	Access_token  string `json:"access_token"`
	ExpireIn      string `json:"expire_in"`
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

//授权请求参数
type VoiceClient struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Authorizer
}

//Authorizer 用于设置access_token
//可以通过RESTFul api的方式从百度方获取
//有效期为一个月，可以存至数据库中然后从数据库中获取
type Authorizer interface {
	Authorize(client *VoiceClient) error
}

type DefaultAuthorizer struct {
}

func (da DefaultAuthorizer) Authorize(client *VoiceClient) error {
	resp, err := req.Post(VOICE_AUTH_URL, req.Param{
		"grant_type":    "client_credentials",
		"client_id":     client.ClientID,
		"client_secret": client.ClientSecret,
	})
	if err != nil {
		return err
	}
	var result AuthResponse
	if err := resp.ToJSON(&result); err != nil {
		client.AccessToken = result.Access_token
	}
	return nil
}

//TextToSpeech 将文字转换为语音
func (vc *VoiceClient) TextToSpeech(txt string) ([]byte, error) {
	vc.auth()
	params := req.Param{
		"tex":  "",             //必填	合成的文本，使用UTF-8编码，请注意文本长度必须小于1024字节
		"lan":  "zh",           //必填	语言选择,目前只有中英文混合模式，填写固定值zh
		"tok":  vc.AccessToken, //必填	开放平台获取到的开发者access_token（见上面的“鉴权认证机制”段落）
		"ctp":  "1",            //必填	客户端类型选择，web端填写固定值1
		"cuid": "random char",  //必填	用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
		"spd":  "5",            //选填	语速，取值0-9，默认为5中语速
		"pit":  "5",            //选填	音调，取值0-9，默认为5中语调
		"vol":  "5",            //选填	音量，取值0-15，默认为5中音量
		"per":  "1",            //选填   发音人选择, 0为普通女声，1为普通男声，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
	}
	resp, err := req.Post(TSN_URL, params)
	if err != nil {
		return []byte{}, nil
	}
	respHeader := resp.Response().Header
	contentType, ok := respHeader["Content-Type"]
	if !ok {
		return []byte{}, errors.New("No Content-Type Set.")
	}
	if contentType[0] == "audio/mp3" {
		respBody, err := ioutil.ReadAll(resp.Response().Body)
		if err != nil {
			return []byte{}, nil
		}
		return respBody, nil
	} else {
		return []byte{}, errors.New(string(resp.ToString()))
	}
}

//SpeechToText 将语音翻译成文字
func (vc *VoiceClient) SpeechToText(path string) ([]string, error) {
	vc.auth()
	return []string{}, nil
}

func (vc *VoiceClient) auth() {
	vc.Authorizer.Authorize(vc)
}

func (vc *VoiceClient) SetAuther(auth Authorizer) {
	vc.Authorizer = auth
}

func NewVoiceClient(ApiKEY, secretKey string) *VoiceClient {
	return &VoiceClient{
		ClientID:     ApiKEY,
		ClientSecret: secretKey,
		Authorizer:   DefaultAuthorizer{},
	}
}
