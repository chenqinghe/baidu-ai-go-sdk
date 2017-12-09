// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package voice

import (
	"errors"

	"io/ioutil"

	"strings"

	"github.com/chenqinghe/baidu-ai-go-sdk"
	"github.com/imroc/req"
)

const (
	TTS_URL = "http://tsn.baidu.com/text2audio"
	ASR_URL = "http://vop.baidu.com/server_api"
)

const (
	B = 1 << 10 * iota
	KB
	MB
	GB
	TB
	PB
)

var (
	ErrTextTooLong = errors.New("The input string is too long")
)

//VoiceClient represent a voice service application.
type VoiceClient struct {
	*gosdk.Client
}

type TTSParams struct {
	Text       string
	Token      string
	Cuid       string
	ClientType int
	Language   string
	Speed      int
	Pitch      int
	Volume     int
	Person     int
}
type TTSParam func(*TTSParams)

func Cuid(str string) TTSParam {
	if len(str) > 60 {
		str = string(str[:60])
	}
	return func(p *TTSParams) {
		p.Cuid = str
	}
}

func Speed(spd int) TTSParam {
	if spd > 9 {
		spd = 9
	}
	if spd < 0 {
		spd = 0
	}
	return func(p *TTSParams) {
		p.Speed = spd
	}
}

func Pitch(pit int) TTSParam {
	if pit > 9 {
		pit = 9
	}
	if pit < 0 {
		pit = 0
	}
	return func(p *TTSParams) {
		p.Pitch = pit
	}
}

func Volume(vol int) TTSParam {
	if vol > 15 {
		vol = 15
	}
	if vol < 0 {
		vol = 0
	}
	return func(p *TTSParams) {
		p.Volume = vol
	}
}

func Person(per int) TTSParam {
	if per != 0 && per != 1 && per != 3 && per != 4 {
		per = 0
	}
	return func(p *TTSParams) {
		p.Person = per
	}
}

//语音识别响应信息
type ASRResponse struct {
	CorpusNo string   `json:"corpus_no"`
	ERRMSG   string   `json:"err_msg"`
	ERRNO    int      `json:"err_no"`
	Result   []string `json:"result"`
	SN       string   `json:"sn"`
}

//语音识别参数
type ASRParams struct {
	Format  string `json:"format"`  //语音的格式，pcm 或者 wav 或者 amr。不区分大小写
	Rate    int    `json:"rate"`    //采样率，支持 8000 或者 16000
	Channel int    `json:"channel"` //声道数，仅支持单声道，请填写固定值 1
	Cuid    string `json:"cuid"`    //用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
	Token   string `json:"token"`   //开放平台获取到的开发者access_token
	Lan     string `json:"lan"`     //语种选择，默认中文（zh）。 中文=zh、粤语=ct、英文=en，不区分大小写
	Speech  string `json:"speech"`  //真实的语音数据 ，需要进行base64 编码。与len参数连一起使用
	Len     int    `json:"len"`     //原始语音长度，单位字节
}

//TextToSpeech 语音合成，将文字转换为语音
func (vc *VoiceClient) TextToSpeech(txt string, params ...TTSParam) ([]byte, error) {

	if len(txt) >= 1024 {
		return nil, ErrTextTooLong
	}
	if err := vc.Auth(); err != nil {
		return nil, err
	}

	ttsparams := &TTSParams{
		Text:       txt,
		Token:      vc.AccessToken,
		Cuid:       "",
		ClientType: 1,
		Language:   "zh",
		Speed:      5,
		Pitch:      5,
		Volume:     5,
		Person:     0,
	}

	for _, param := range params {
		param(ttsparams)
	}

	resp, err := req.Post(TTS_URL, req.Param(gosdk.StructToMap(params)))
	if err != nil {
		return nil, err
	}
	respHeader := resp.Response().Header
	contentType, ok := respHeader["Content-Type"]
	if !ok {
		return nil, errors.New("No Content-Type Set.")
	}
	if contentType[0] == "audio/mp3" {
		respBody, err := ioutil.ReadAll(resp.Response().Body)
		if err != nil {
			return nil, err
		}
		return respBody, nil
	} else {
		respStr, err := resp.ToString()
		if err != nil {
			return nil, err
		}
		return nil, errors.New("调用服务失败：" + respStr)
	}

}

//SpeechToText 语音识别，将语音翻译成文字
func (vc *VoiceClient) SpeechToText(ap ASRParams) ([]string, error) {
	if ap.Len > 8*10*MB {
		return []string{}, errors.New("文件大小不能超过10M")
	}
	if err := vc.Auth(); err != nil {
		return []string{}, err
	}
	ap.Token = vc.AccessToken
	resp, err := req.Post(ASR_URL, req.Header{
		"Content-Type": "application/json",
	}, req.BodyJSON(ap))
	if err != nil {
		return []string{}, err
	}
	var rs ASRResponse
	if err := resp.ToJSON(&rs); err != nil {
		return []string{}, err
	}
	if !strings.Contains(rs.ERRMSG, "success") || rs.ERRNO != 0 {
		return []string{}, errors.New("调用服务失败：" + rs.ERRMSG)
	}
	return rs.Result, nil
}

func NewVoiceClient(ApiKey, secretKey string) *VoiceClient {
	return &VoiceClient{
		Client:    gosdk.NewClient(ApiKey, secretKey),
		TTSConfig: defaultTTSConfig,
	}
}
