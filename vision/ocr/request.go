package ocr

import "github.com/imroc/req"

type OCRResponse struct {
	*req.Resp
}

func (oc *OCRClient) doRequest(url string, params map[string]interface{}) (response *OCRResponse, err error) {

	if err := oc.Auth(); err != nil {
		return nil, err
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	url += "?access_token=" + oc.AccessToken

	resp, err := req.Post(url, req.Param(params), header)
	if err != nil {
		return nil, err
	}
	return &OCRResponse{
		Resp: resp,
	}, nil
}
