package ocr

import (
	"encoding/base64"
	"github.com/chenqinghe/baidu-ai-go-sdk"
	"io"
	"io/ioutil"
)

const (
	OCR_GENERAL_BASIC_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	OCR_GENERAL_WITH_LOCATION_URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	OCR_GENERAL_ENHANCED_URL      = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"
)

type OCRClient struct {
	*gosdk.Client
}

func NewOCRClient(apiKey, secretKey string) *OCRClient {
	return &OCRClient{
		Client: gosdk.NewClient(apiKey, secretKey),
	}
}

//GeneralRecognizeBasic 通用文字识别
//img 图片二进制数据
//conf 请求参数
func (oc *OCRClient) GeneralRecognizeBasic(imageReader io.Reader, params ...RequestParam) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}

	imgBytes, err := ioutil.ReadAll(imageReader)
	if err != nil {
		return nil, err
	}

	encodedImgStr := base64.StdEncoding.EncodeToString(imgBytes)

	conf := defaultGeneralBasicParams
	conf["image"] = encodedImgStr

	for _, fn := range params {
		fn(conf)
	}

	var url = OCR_GENERAL_BASIC_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)
}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(imageReader io.Reader, params ...RequestParam) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}

	imgBytes, err := ioutil.ReadAll(imageReader)
	if err != nil {
		return nil, err
	}

	encodedImgStr := base64.StdEncoding.EncodeToString(imgBytes)
	conf := defaultGeneralWithLocationParams
	conf["image"] = encodedImgStr

	for _, fn := range params {
		fn(conf)
	}

	var url = OCR_GENERAL_WITH_LOCATION_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
func (oc *OCRClient) GeneralRecognizeEnhanced(imageReader io.Reader, params ...RequestParam) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}

	imgBytes, err := ioutil.ReadAll(imageReader)
	if err != nil {
		return nil, err
	}

	encodedImgStr := base64.StdEncoding.EncodeToString(imgBytes)

	conf := defaultDeneralEnhancedParams
	conf["image"] = encodedImgStr

	for _, fn := range params {
		fn(conf)
	}

	url := OCR_GENERAL_ENHANCED_URL + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)

}
