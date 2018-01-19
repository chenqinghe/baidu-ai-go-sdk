package face

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
)

const (
	FACE_MATCH_URL = "https://aip.baidubce.com/rest/2.0/face/v2/match"
)

func (fc *FaceClient) Match(img1, img2 *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}

	imgContent1, err := img1.Base64()
	imgContent2, err := img2.Base64()
	if err != nil { //任意一个出错都不行
		return nil, err
	}

	options["image"] = fmt.Sprintf("%s,%s", imgContent1, imgContent2)

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	url := FACE_MATCH_URL + "?=" + fc.AccessToken
	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
