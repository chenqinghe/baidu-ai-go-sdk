package ocr

import "github.com/imroc/req"

func doRequest(url string, params map[string]interface{}) (rs []byte, err error) {

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(params), header)
	if err != nil {
		return
	}
	return resp.ToBytes()
}