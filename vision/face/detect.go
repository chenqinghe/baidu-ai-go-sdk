package face

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
)

const (
	FACE_DETECT_URL = "https://aip.baidubce.com/rest/2.0/face/v2/detect"
)

type FaceResponse struct {
	*req.Resp
}

func (fc *FaceClient) DetectAndAnalysis(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {

	if err := fc.Auth(); err != nil {
		return nil, err
	}

	url := FACE_DETECT_URL + "?access_token=" + fc.AccessToken

	base64Str, err := image.Base64()
	if err != nil {
		return nil, err
	}
	options["image"] = base64Str

	resp, err := req.Post(url, req.Param(options))
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
