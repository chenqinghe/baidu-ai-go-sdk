package v2

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
	"strings"
)

const (
	faceIdentifyUrl      = "https://aip.baidubce.com/rest/2.0/face/v2/identify"
	faceVerifyUrl        = "https://aip.baidubce.com/rest/2.0/face/v2/verify"
	faceMultiIdentifyUrl = "https://aip.baidubce.com/rest/2.0/face/v2/multi-identify"
)

type IdentifyResponse struct {
	*req.Resp
}

func (fc FaceClient) Identify(img *vision.Image, groupID []string, options map[string]interface{}) (*IdentifyResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	encodedImg, err := img.Base64Encode()
	if err != nil {
		return nil, err
	}
	options["image"] = encodedImg
	options["group_id"] = strings.Join(groupID, ",")

	url := fmt.Sprintf("%s?access_token=%s", faceIdentifyUrl, fc.AccessToken)
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &IdentifyResponse{resp}, nil

}

type VerifyResponse struct {
	*req.Resp
}

func (fc FaceClient) Verify(img *vision.Image, uid string, groupID string, options map[string]interface{}) (*VerifyResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	encodedImg, err := img.Base64Encode()
	if err != nil {
		return nil, err
	}
	options["image"] = encodedImg
	options["uid"] = uid
	options["group_id"] = groupID

	url := fmt.Sprintf("%s?access_token=%s", faceVerifyUrl, fc.AccessToken)
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}
	return &VerifyResponse{resp}, nil
}

type MultiVerifyResponse struct {
	*req.Resp
}

func (fc FaceClient) MultiVerify(image vision.Image, groupIDs []string, options map[string]interface{}) (*MultiVerifyResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}

	encodedImage, err := image.Base64Encode()
	if err != nil {
		return nil, err
	}
	options["image"] = encodedImage
	options["group_id"] = strings.Join(groupIDs, ",")

	url := fmt.Sprintf("%s?access_token=%s", faceMultiIdentifyUrl, fc.AccessToken)
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}
	return &MultiVerifyResponse{resp}, nil

}
