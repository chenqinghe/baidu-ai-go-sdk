package v3

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

const (
	facePersonVerifyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/person/verify"
)

type PersonVerifyResponse struct {
	*req.Resp
}

func (fc FaceClient) PersonVerify(image, imageType, idCardNumber, name string, options map[string]string) (*PersonVerifyResponse, error) {
	if imageType != "BASE64" && imageType != "URL" && imageType != "FACE_TOKEN" {
		return nil, errors.New("image_type is invalid")
	}
	if imageType == "BASE64" && len(image) > 2<<(20-1)*2 {
		return nil, errors.New("image length is invalid")
	}
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	options["image"] = image
	options["image_type"] = imageType
	options["id_card_number"] = idCardNumber
	options["name"] = name

	url := fmt.Sprintf("%s?access_token=%s", facePersonVerifyUrl, fc.AccessToken)
	resp, err := req.Post(url, req.BodyJSON(&options))
	if err != nil {
		return nil, err
	}
	return &PersonVerifyResponse{resp}, nil
}
