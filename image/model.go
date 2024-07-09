package image

import (
	"encoding/base64"
	"io"
)

// Input 三选一
type Input struct {
	ImageUrl    string    `json:"image_url"`
	File        io.Reader `json:"file"`
	ImageBase64 string    `json:"image_base_64"`
}

func (i *Input) encode(maxSize int) (map[string]interface{}, error) {
	var invokeParam = make(map[string]interface{})
	switch {
	case i.File != nil:
		binary, err := io.ReadAll(i.File)
		if err != nil {
			return nil, err
		}
		base64Encode := base64.StdEncoding.EncodeToString(binary)
		if len(base64Encode) > maxSize {
			return nil, ErrImageTooLarge
		}
		invokeParam["image"] = base64Encode
	case i.ImageBase64 != "":
		if len(i.ImageBase64) > maxSize {
			return nil, ErrImageTooLarge
		}
		invokeParam["image"] = i.ImageBase64
	case i.ImageUrl != "":
		invokeParam["url"] = i.ImageUrl
	default:
		return nil, ErrInvalidImage
	}
	return invokeParam, nil
}

type EnhanceResponse struct {
	Image string `json:"image"`
	LogID int    `json:"log_id"`
}

type generalResponse struct {
	EnhanceResponse
	APIError
}

func (g *generalResponse) success() bool {
	return g.ErrorCode == 0
}
