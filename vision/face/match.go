package face

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
	url2 "net/url"
	"strings"
)

const (
	FACE_MATCH_URL = "https://aip.baidubce.com/rest/2.0/face/v2/match"
)

func (fc *FaceClient) Match(img1, img2 *vision.Image, options map[string]string) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}

	imgContent1, err := img1.Base64Encode()
	imgContent2, err := img2.Base64Encode()
	if err != nil { //任意一个出错都不行
		return nil, err
	}

	options["images"] = fmt.Sprintf("%s,%s", url2.QueryEscape(imgContent1), url2.QueryEscape(imgContent2))

	body := ""
	for k, v := range options {
		body += fmt.Sprintf("%s=%s&", k, v)
	}
	body = strings.TrimRight(body, "&")

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	url := FACE_MATCH_URL + "?access_token=" + fc.AccessToken
	resp, err := req.Post(url, header, body)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
