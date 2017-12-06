package ocr

import (
	"encoding/base64"

	sdk "github.com/chenqinghe/baidu-ai-go-sdk/internal"
	"github.com/imroc/req"
)

const (
	OCR_GENERAL_BASIC_URL         string = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	OCR_GENERAL_WITH_LOCATION_URL        = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	OCR_GENERAL_ENHANCED_URL             = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"
)

var defaultGeneralBasicParams = map[string]string{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
}

var defaultGeneralWithLocationParams = map[string]string{
	"image":                 "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"recognize_granularity": "big",     //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"language_type":         "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction":      "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
	"detect_language":       "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"vertexes_location":     "false",   //是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
	"probability":           "false",   //是否返回识别结果中每一行的置信度
}

var defaultDeneralEnhancedParams = map[string]string{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"probability":      "false",   //是否返回识别结果中每一行的置信度
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
		return nil, err
	}
	encodedImgStr := base64.StdEncoding.EncodeToString(img)
	conf["image"] = encodedImgStr

	conf = parseParams(defaultGeneralBasicParams, conf)

	var url string = OCR_GENERAL_BASIC_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)
}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(img []byte, conf map[string]string) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}
	encodedImgStr := base64.StdEncoding.EncodeToString(img)
	conf["image"] = encodedImgStr
	conf = parseParams(defaultGeneralWithLocationParams, conf)

	var url string = OCR_GENERAL_WITH_LOCATION_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
func (oc *OCRClient) GeneralRecognizeEnhanced(img []byte, conf map[string]string) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}
	encodedImgStr := base64.StdEncoding.EncodeToString(img)
	conf["image"] = encodedImgStr

	conf = parseParams(defaultDeneralEnhancedParams, conf)

	url := OCR_GENERAL_ENHANCED_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)

}

func parseParams(def, need map[string]string) map[string]string {
	for key, _ := range def {
		if val, ok := need[key]; ok {
			def[key] = val
		}
	}
	return def
}

func doRequest(url string, params map[string]interface{}) (rs []byte, err error) {

	resp, err := req.Post(url, req.Param(params), req.Header{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return
	}
	rs, err = resp.ToBytes()
	return

}
