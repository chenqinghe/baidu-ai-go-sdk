package image

import (
	"errors"
	"fmt"
)

var (
	ErrImageTooLarge = errors.New("image to large")
	ErrInvalidImage  = errors.New("invalid image")
)

// APIError https://ai.baidu.com/ai-doc/IMAGEPROCESS/Ek3bclpgv
type APIError struct {
	ErrorCode  int    `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
	StatusCode string `json:"-"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%#v", e)
}
