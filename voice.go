// 语音处理
// 利用百度RESTFul API 进行语音及文字的相互转换
package dueros

import "github.com/imroc/req"

const TSN_URL string = "http://tsn.baidu.com/text2audio"
const VOP_URL string = "http://vop.baidu.com/server_api"

type TSNResponse struct {
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

//TextToSpeech 将文字转换为语音
func TextToSpeech() ([]string, error) {
	header := req.Header{}
	resp, err := req.Post(TSN_URL, header)
	if err != nil {
		return []string{}, nil
	}
	var result TSNResponse
	if err := resp.ToJSON(&result); err != nil {
		return []string{}, err
	}
	return ,nil
}

//SpeechToText 讲语音翻译成文字
func SpeechToText() {

}

func auth() {

}
