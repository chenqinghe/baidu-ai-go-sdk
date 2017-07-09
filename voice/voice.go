// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package voice

import (
	"errors"

	"io/ioutil"

	sdk "github.com/chenqinghe/baidu-ai-go-sdk/internal"
	"github.com/imroc/req"
	"net"
	"strconv"
)

const (
	TTS_URL string = "http://tsn.baidu.com/text2audio"
	ASR_URL string = "http://vop.baidu.com/server_api"
)
const (
	B int = 1 << (10 * iota)
	KB
	MB
)

var ErrNoTTSConfig = errors.New("No TTSConfig.please set TTSConfig correctlly first or call method UseDefaultTTSConfig")
var ErrTextTooLong = errors.New("The input string is too long")

//VoiceClient 代表一个语音服务应用
type VoiceClient struct {
	*sdk.Client
	TTSConfig *TTSConfig
}

//语音合成参数
type TTSConfig struct {
	SPD int //语速，取值0-9，默认为5中语速
	PIT int //音调，取值0-9，默认为5中语调
	VOL int //音量，取值0-15，默认为5中音量
	PER int //发音人选择, 0为普通女声，1为普通男声，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
}

var defaultTTSConfig = &TTSConfig{
	SPD: 5,
	PIT: 5,
	VOL: 5,
	PER: 0,
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

func (vc *VoiceClient) UseDefaultTTSConfig() *VoiceClient {
	vc.TTSConfig = defaultTTSConfig
	return vc
}

//TextToSpeech 语音合成，将文字转换为语音
func (vc *VoiceClient) TextToSpeech(txt string) ([]byte, error) {
	if len(txt) >= 1024 {
		return []byte{}, ErrTextTooLong
	}
	if err := vc.Auth(); err != nil {
		return []byte{}, err
	}
	if vc.TTSConfig == nil {
		return []byte{}, ErrNoTTSConfig
	}
	itfcs, err := net.Interfaces()
	if err != nil {
		return []byte{}, err
	}
	mac := itfcs[0].HardwareAddr.String()
	params := req.Param{
		"tex":  txt,                            //必填	合成的文本，使用UTF-8编码，请注意文本长度必须小于1024字节
		"lan":  "zh",                           //必填	语言选择,目前只有中英文混合模式，填写固定值zh
		"tok":  vc.AccessToken,                 //必填	开放平台获取到的开发者access_token（见上面的“鉴权认证机制”段落）
		"ctp":  "1",                            //必填	客户端类型选择，web端填写固定值1
		"cuid": mac,                            //必填	用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
		"spd":  strconv.Itoa(vc.TTSConfig.SPD), //选填	语速，取值0-9，默认为5中语速
		"pit":  strconv.Itoa(vc.TTSConfig.PIT), //选填	音调，取值0-9，默认为5中语调
		"vol":  strconv.Itoa(vc.TTSConfig.VOL), //选填	音量，取值0-15，默认为5中音量
		"per":  strconv.Itoa(vc.TTSConfig.PER), //选填 发音人选择, 0为普通女声，1为普通男声，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
	}
	resp, err := req.Post(TTS_URL, params)
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
		respStr, err := resp.ToString()
		if err != nil {
			return []byte{}, err
		}
		return []byte{}, errors.New("调用服务失败：" + respStr)
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
	if rs.ERRMSG != "success." || rs.ERRNO != 0 {
		return []string{}, errors.New("调用服务失败：" + rs.ERRMSG)
	}
	return rs.Result, nil
}

func NewVoiceClient(ApiKey, secretKey string) *VoiceClient {
	return &VoiceClient{
		Client: sdk.NewClient(ApiKey, secretKey),
	}
}
