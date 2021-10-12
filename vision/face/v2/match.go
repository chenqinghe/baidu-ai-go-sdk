package v2

import (
	"fmt"

	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
)

const (
	faceMatchUrl = "https://aip.baidubce.com/rest/2.0/face/v2/match"
)

func (fc FaceClient) Match(img1, img2 *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}

	imgContent1, err := img1.Base64Encode()
	if err != nil {
		return nil, err
	}
	imgContent2, err := img2.Base64Encode()
	if err != nil {
		return nil, err
	}

	options["images"] = fmt.Sprintf("%s,%s", imgContent1, imgContent2)

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	url := fmt.Sprintf("%s?access_token=%s", faceMatchUrl, fc.AccessToken)
	resp, err := req.Post(url, header, req.Param(options))
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
