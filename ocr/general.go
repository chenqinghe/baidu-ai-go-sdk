package ocr

import (
	"encoding/base64"
	sdk "github.com/chenqinghe/baidu-ai-go-sdk/internal"
	"github.com/imroc/req"
)

const OCR_GENERAL_BASIC_URL string = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

var defaultGeneralBasicParams = map[string]string{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
}

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
	if err := oc.Auth(); err != nil {
		return []byte{}, err
	}
	encodedImgStr := base64.StdEncoding.EncodeToString(img)
	defaultGeneralBasicParams["image"] = encodedImgStr
	for key, _ := range defaultGeneralBasicParams {
		if v, ok := conf[key]; ok {
			defaultGeneralBasicParams[key] = v
		}
	}
	var url string = OCR_GENERAL_BASIC_URL + "?access_token=" + oc.AccessToken
	resp, err := req.Post(url, req.Param(defaultGeneralBasicParams), req.Header{
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
