// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package voice

import (
	"errors"

	"io/ioutil"

	"net"

	"encoding/json"
	"github.com/imroc/req"
)

const TTS_URL = "http://tsn.baidu.com/text2audio"

var (
	ErrTextTooLong = errors.New("The input string is too long")
)

type TTSParams struct {
	Text       string `json:"tex"`
	Token      string `json:"tok"`
	Cuid       string `json:"cuid"`
	ClientType int    `json:"ctp"`
	Language   string `json:"lan"`
	Speed      int    `json:"spd"`
	Pitch      int    `json:"pit"`
	Volume     int    `json:"vol"`
	Person     int    `json:"per"`
}

type TTSParam func(params *TTSParams)

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

	var cuid string
	netitfs, err := net.Interfaces()
	if err != nil {
		cuid = "anonymous"
	} else {
		cuid = netitfs[0].HardwareAddr.String()
	}

	ttsparams := &TTSParams{
		Text:       txt,
		Token:      vc.AccessToken,
		Cuid:       cuid,
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

	t, err := json.Marshal(ttsparams)
	if err != nil {
		return nil, errors.New("serialize failed: " + err.Error())
	}
	var p  = req.Param{}
	if err := json.Unmarshal(t, &p); err != nil {
		return nil, err
	}

	resp, err := req.Post(TTS_URL, p)
	if err != nil {
		return nil, err
	}
	
	//通过Content-Type的头部来确定是否服务端合成成功。
	//http://ai.baidu.com/docs#/TTS-API/top
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
