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

const (
	OCR_WEBIMAGE_URL       = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"
	OCR_IDCARD_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	OCR_BANKCARD_URL       = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	OCR_DRIVERLICENSE_URL  = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
	OCR_VEHICLELICENSE_URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"
	OCR_LICENSEPLATE_URL   = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	OCR_FORM_URL           = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"
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

	return oc.ocr(imageReader, OCR_GENERAL_BASIC_URL, defaultGeneralBasicParams, params...)

}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_GENERAL_WITH_LOCATION_URL, defaultGeneralWithLocationParams, params...)

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
func (oc *OCRClient) GeneralRecognizeEnhanced(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_GENERAL_ENHANCED_URL, defaultDeneralEnhancedParams, params...)

}

func (oc *OCRClient) WebImageRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_WEBIMAGE_URL, defaultWebimgParams, params...)

}

func (oc *OCRClient) IdcardRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_IDCARD_URL, defaultIdcardParams, params...)

}

func (oc *OCRClient) BankcardRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_BANKCARD_URL, defaultBankcardParams, params...)

}

func (oc *OCRClient) DriverLicenseRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_DRIVERLICENSE_URL, defaultDriverLicenseParams, params...)

}

func (oc *OCRClient) VehicleLicenseRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_VEHICLELICENSE_URL, defaultVehicleLicenseParams, params...)

}

func (oc *OCRClient) LicensePlateRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_LICENSEPLATE_URL, defaultLicensePlateParams, params...)

}

func (oc *OCRClient) FromdataRecognize(imageReader io.Reader, params ...RequestParam) ([]byte, error) {

	return oc.ocr(imageReader, OCR_FORM_URL, defaultFormParams, params...)

}

func (oc *OCRClient) ocr(imageReader io.Reader, url string, def map[string]interface{}, params ...RequestParam) ([]byte, error) {
	requestParams, err := parseRequestParam(imageReader, def, params...)
	if err != nil {
		return nil, err
	}

	return oc.doRequest(url, requestParams)
}

func parseRequestParam(imageReader io.Reader, def map[string]interface{}, params ...RequestParam) (map[string]interface{}, error) {

	imageBytes, err := ioutil.ReadAll(imageReader)
	if err != nil {
		return nil, err
	}
	imageBase64Str := base64.StdEncoding.EncodeToString(imageBytes)

	def["image"] = imageBase64Str

	for _, fn := range params {
		fn(def)
	}

	return def, nil

}
