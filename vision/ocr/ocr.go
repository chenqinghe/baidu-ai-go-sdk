package ocr

import (
	"errors"
	"github.com/chenqinghe/baidu-ai-go-sdk"
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
func (oc *OCRClient) GeneralRecognizeBasic(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_GENERAL_BASIC_URL, defaultGeneralBasicParams, params...)

}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_GENERAL_WITH_LOCATION_URL, defaultGeneralWithLocationParams, params...)

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
func (oc *OCRClient) GeneralRecognizeEnhanced(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_GENERAL_ENHANCED_URL, defaultDeneralEnhancedParams, params...)

}

func (oc *OCRClient) WebImageRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_WEBIMAGE_URL, defaultWebimgParams, params...)

}

func (oc *OCRClient) IdcardRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_IDCARD_URL, defaultIdcardParams, params...)

}

func (oc *OCRClient) BankcardRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_BANKCARD_URL, defaultBankcardParams, params...)

}

func (oc *OCRClient) DriverLicenseRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_DRIVERLICENSE_URL, defaultDriverLicenseParams, params...)

}

func (oc *OCRClient) VehicleLicenseRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_VEHICLELICENSE_URL, defaultVehicleLicenseParams, params...)

}

func (oc *OCRClient) LicensePlateRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_LICENSEPLATE_URL, defaultLicensePlateParams, params...)

}

func (oc *OCRClient) FromdataRecognize(image *Image, params ...RequestParam) ([]byte, error) {

	return oc.ocr(image, OCR_FORM_URL, defaultFormParams, params...)

}

func (oc *OCRClient) ocr(image *Image, url string, def map[string]interface{}, params ...RequestParam) ([]byte, error) {
	requestParams, err := parseRequestParam(image, def, params...)
	if err != nil {
		return nil, err
	}

	return oc.doRequest(url, requestParams)
}

func parseRequestParam(image *Image, def map[string]interface{}, params ...RequestParam) (map[string]interface{}, error) {

	if image.Reader == nil {
		if image.Url == "" {
			return nil, errors.New("image source is empty")
		} else {
			def["url"] = image.Url
		}
	} else {
		base64Str, err := image.Base64()
		if err != nil {
			return nil, err
		}
		def["image"] = base64Str
	}

	for _, fn := range params {
		fn(def)
	}

	return def, nil

}
