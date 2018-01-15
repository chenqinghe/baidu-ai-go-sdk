package voice

import (
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"

	"net"

	"fmt"
	"github.com/imroc/req"
)

const ASR_URL = "http://vop.baidu.com/server_api"

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
	Format   string `json:"format"`  //语音的格式，pcm 或者 wav 或者 amr。不区分大小写
	Rate     int    `json:"rate"`    //采样率，支持 8000 或者 16000
	Channel  int    `json:"channel"` //声道数，仅支持单声道，请填写固定值 1
	Cuid     string `json:"cuid"`    //用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
	Token    string `json:"token"`   //开放平台获取到的开发者access_token
	Language string `json:"lan"`     //语种选择，默认中文（zh）。 中文=zh、粤语=ct、英文=en，不区分大小写
	Speech   string `json:"speech"`  //真实的语音数据 ，需要进行base64 编码。与len参数连一起使用
	Length   int    `json:"len"`     //原始语音长度，单位字节
}

type ASRParam func(params *ASRParams)

func Format(fmt string) ASRParam {

	if fmt != "pcm" && fmt != "wav" && fmt != "amr" {
		fmt = "pcm"
	}
	return func(params *ASRParams) {
		params.Format = fmt
	}
}

func Rate(rate int) ASRParam {
	if rate != 8000 && rate != 16000 {
		rate = 8000
	}
	return func(params *ASRParams) {
		params.Rate = rate
	}
}

func Channel(c int) ASRParam {
	return func(params *ASRParams) {
		params.Channel = 1 //固定值1
	}
}

func Language(lang string) ASRParam {
	if lang != "zh" && lang != "ct" && lang != "en" {
		lang = "zh"
	}
	return func(params *ASRParams) {
		params.Language = lang
	}
}

////SpeechToText 语音识别，将语音翻译成文字
func (vc *VoiceClient) SpeechToText(reader io.Reader, params ...ASRParam) ([]string, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if len(content) > 10*MB {
		return nil, errors.New("文件大小不能超过10M")
	}

	spch := base64.StdEncoding.EncodeToString(content)

	var cuid string
	netitfs, err := net.Interfaces()
	if err != nil {
		cuid = "anonymous"
	} else {
		cuid = netitfs[0].HardwareAddr.String()
	}

	asrParams := &ASRParams{
		Format:   "pcm",
		Rate:     8000,
		Channel:  1,
		Cuid:     cuid,
		Token:    vc.AccessToken,
		Language: "zh",
		Speech:   spch,
		Length:   len(content),
	}

	for _, fn := range params {
		fn(asrParams)
	}

	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Post(ASR_URL, header, req.BodyJSON(asrParams))
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	var asrResponse *ASRResponse
	if err := resp.ToJSON(asrResponse); err != nil {
		return nil, err
	}

	if asrResponse.ERRNO != 0 {
		return nil, errors.New("调用服务失败：" + asrResponse.ERRMSG)
	}

	return asrResponse.Result, nil

}
