package ocr

import (
	"encoding/base64"
	sdk "github.com/chenqinghe/baidu-ai-go-sdk/internal"
	"github.com/imroc/req"
)

const OCR_GENERAL_BASIC_URL string = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

type OCRClient struct {
	*sdk.Client
}

func NewOCRClient(apiKey, secretKey string) *OCRClient {
	return &OCRClient{
		Client: sdk.NewClient(apiKey, secretKey),
	}
}

//GeneralRecognizeBasic 通用文字识别
//img 图片二进制数据
//conf 请求参数
func (oc *OCRClient) GeneralRecognizeBasic(img []byte, conf map[string]string) ([]byte, error) {
	oc.Auth()
	encodedImgStr := base64.StdEncoding.EncodeToString(img)
	var params map[string]string = map[string]string{
		"image":            encodedImgStr,
		"language_type":    "CHN_ENG",
		"detect_direction": "false",
		"detect_language":  "false",
	}
	for key, _ := range params {
		if v, ok := conf[key]; ok {
			params[key] = v
		}
	}
	var url string = OCR_GENERAL_BASIC_URL + "?access_token=" + oc.AccessToken
	resp, err := req.Post(url, req.Param(params), req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	if err != nil {
		return []byte{}, err
	}
	respByte, err := resp.ToBytes()
	if err != nil {
		return []byte{}, err
	}
	return respByte, nil
}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(img []byte, conf map[string]string) {

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
func (oc *OCRClient) GeneralRecognizeEnhanced() {

}
