package v3

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

const (
	faceMatchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/match"
)

type FaceMatchResponse struct {
	*req.Resp
}

type FaceMatchImage struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	FaceType        string `json:"face_type,omitempty"`        // 默认LIVE
	QualityControl  string `json:"quality_control,omitempty"`  // 默认 NONE
	LivenessControl string `json:"liveness_control,omitempty"` // 默认 NONE
	FaceSortType    string `json:"face_sort_type,omitempty"`   // 默认为0
}

func (fc FaceClient) FaceMatch(images [2]FaceMatchImage) (*FaceMatchResponse, error) {
	for _, image := range images {
		if image.ImageType != "BASE64" && image.ImageType != "URL" && image.ImageType != "FACE_TOKEN" {
			return nil, errors.New("image_type is invalid")
		}
		if image.ImageType == "BASE64" && len(image.Image) > 2<<(20-1)*2 {
			return nil, errors.New("image length is invalid")
		}
	}
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", faceMatchUrl, fc.AccessToken)
	resp, err := req.Post(url, req.BodyJSON(&images))
	if err != nil {
		return nil, err
	}
	return &FaceMatchResponse{resp}, nil
}
