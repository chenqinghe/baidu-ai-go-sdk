package v2

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

const (
	facePersonVerifyUrl = "https://aip.baidubce.com/rest/2.0/face/v2/person/verify"
)

type PersonVerifyResponse struct {
	*req.Resp
}

/// image cannot exceed 5M
func (fc FaceClient) PersonVerify(image, idCardNumber, name string, options map[string]string) (*PersonVerifyResponse, error) {
	if len(image) > 2<<(20-1)*5 {
		return nil, errors.New("image length is invalid")
	}

	if fc.AccessToken == "" {
		if err := fc.Auth(); err != nil {
			return nil, err
		}
	}
	options["image"] = image
	options["id_card_number"] = idCardNumber
	options["name"] = name

	url := fmt.Sprintf("%s?access_token=%s", facePersonVerifyUrl, fc.AccessToken)
	resp, err := req.Post(url, req.BodyJSON(&options))
	if err != nil {
		return nil, err
	}
	return &PersonVerifyResponse{resp}, nil
}
