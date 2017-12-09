package voice

import (
	"errors"
	"github.com/imroc/req"
	"strings"
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
	Format  string `json:"format"`  //语音的格式，pcm 或者 wav 或者 amr。不区分大小写
	Rate    int    `json:"rate"`    //采样率，支持 8000 或者 16000
	Channel int    `json:"channel"` //声道数，仅支持单声道，请填写固定值 1
	Cuid    string `json:"cuid"`    //用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
	Token   string `json:"token"`   //开放平台获取到的开发者access_token
	Lan     string `json:"lan"`     //语种选择，默认中文（zh）。 中文=zh、粤语=ct、英文=en，不区分大小写
	Speech  string `json:"speech"`  //真实的语音数据 ，需要进行base64 编码。与len参数连一起使用
	Len     int    `json:"len"`     //原始语音长度，单位字节
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
