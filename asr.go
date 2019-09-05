package gosdk

import (
	"io"
	"net/http"
)

const asrUrl = "http://vop.baidu.com/server_api"

// 语音识别响应信息
type ASRResponse struct {
	CorpusNo string   `json:"corpus_no"`
	ErrMsg   string   `json:"err_msg"`
	ErrNo    int      `json:"err_no"`
	Result   []string `json:"result"`
	SN       string   `json:"sn"`
}

func (c *Client) SpeechToText(reader io.Reader, params Params) ([]string, error) {
	req:= http.NewRequest()
	c.httpclient.Post(asrUrl)
}
