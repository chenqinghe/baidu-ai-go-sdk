// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package voice

import (
	"errors"

	"io/ioutil"

	"github.com/chenqinghe/baidu-ai-go-sdk"
	"github.com/imroc/req"
)

const TTS_URL = "http://tsn.baidu.com/text2audio"

var (
	ErrTextTooLong = errors.New("The input string is too long")
)

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
		Cuid:       "as",
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

	resp, err := req.Post(TTS_URL, req.Param(gosdk.StructToMap(*ttsparams)))
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
