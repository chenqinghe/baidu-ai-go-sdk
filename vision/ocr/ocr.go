package ocr

import (
	"errors"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
)

const (
	OCR_GENERAL_BASIC_URL          = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	OCR_ACCURATE_BASIC_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
	OCR_GENERAL_WITH_LOCATION_URL  = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	OCR_GENERAL_ENHANCED_URL       = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"
	OCR_WEBIMAGE_URL               = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"
	OCR_IDCARD_URL                 = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	OCR_BANKCARD_URL               = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	OCR_DRIVERLICENSE_URL          = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
	OCR_VEHICLELICENSE_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"
	OCR_LICENSEPLATE_URL           = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	OCR_FORM_URL                   = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"
	OCR_VAT_INVOICE_URL            = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	OCR_IOCR_RECOGNISE_URL         = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise"
	OCR_IOCR_RECOGNISE_FINANCE_URL = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise/finance"
)

//GeneralRecognizeBasic 通用文字识别
//识别图片中的文字信息
func (oc *OCRClient) GeneralRecognizeBasic(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_GENERAL_BASIC_URL, defaultGeneralBasicParams, params...)

}

//AccurateRecognizeBasic 通用文字识别(高精度版)
//识别图片中的文字信息
func (oc *OCRClient) AccurateRecognizeBasic(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_ACCURATE_BASIC_URL, defaultAccurateBasicParams, params...)

}

//GeneralRecognizeWithLocation 通用文字识别（含位置信息）
//识别图片中的文字信息（包含文字区域的坐标信息）
func (oc *OCRClient) GeneralRecognizeWithLocation(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_GENERAL_WITH_LOCATION_URL, defaultGeneralWithLocationParams, params...)

}

//GeneralRecognizeEnhanced 通用文字识别（含生僻字）
//识别图片中的文字信息（包含对常见字和生僻字的识别）
func (oc *OCRClient) GeneralRecognizeEnhanced(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_GENERAL_ENHANCED_URL, defaultDeneralEnhancedParams, params...)

}

//WebImageRecognize 网络图片识别
//识别一些网络上背景复杂，特殊字体的文字
func (oc *OCRClient) WebImageRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_WEBIMAGE_URL, defaultWebimgParams, params...)

}

//IdCardRecognize 身份证识别
//识别身份证正反面的文字信息
func (oc *OCRClient) IdCardRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_IDCARD_URL, defaultIdcardParams, params...)

}

//BankcardRecognize 银行卡识别
//识别银行卡的卡号并返回发卡行和卡片性质信息
func (oc *OCRClient) BankcardRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_BANKCARD_URL, defaultBankcardParams, params...)

}

//DriverLicenseRecognize 驾驶证识别
//识别机动车驾驶证所有关键字段
func (oc *OCRClient) DriverLicenseRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_DRIVERLICENSE_URL, defaultDriverLicenseParams, params...)

}

// VehicleLicenseRecognize 行驶证识别
// 识别机动车行驶证所有关键字段
// 默认使用高精度服务，可选快速服务ocr.Accuracy("normal")
func (oc *OCRClient) VehicleLicenseRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_VEHICLELICENSE_URL, defaultVehicleLicenseParams, params...)

}

//LicensePlateRecognize 车牌识别
//对小客车的车牌进行识别
func (oc *OCRClient) LicensePlateRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_LICENSEPLATE_URL, defaultLicensePlateParams, params...)

}

//FormDataRecognize 表格文字识别
//自动识别表格线及表格内容，结构化输出表头、表尾及每个单元格的文字内容
func (oc *OCRClient) FormDataRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {

	return oc.ocr(image, OCR_FORM_URL, defaultFormParams, params...)

}

//VATInvoiceRecognize 增值税发票识别
func (oc *OCRClient) VATInvoiceRecognize(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_VAT_INVOICE_URL, defaultVATInvoiceParams, params...)

}

//TODO:营业执照识别

//TODO:通用票据识别

//IocrRecognise 自定义模板文字识别
func (oc *OCRClient) IocrRecognise(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_IOCR_RECOGNISE_URL, defaultIocrRecogniseParams, params...)
}

//IocrRecogniseFinance 自定义模板文字识别  财会版
func (oc *OCRClient) IocrRecogniseFinance(image *vision.Image, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_IOCR_RECOGNISE_FINANCE_URL, defaultIocrRecogniseFinanceParams, params...)
}

func (oc *OCRClient) ocr(image *vision.Image, url string, def map[string]interface{}, params ...RequestParam) (*OCRResponse, error) {
	requestParams, err := parseRequestParam(image, def, params...)
	if err != nil {
		return nil, err
	}

	return oc.doRequest(url, requestParams)
}

func parseRequestParam(image *vision.Image, def map[string]interface{}, params ...RequestParam) (map[string]interface{}, error) {
	if image.Reader == nil {
		if image.Url == "" {
			return nil, errors.New("image source is empty")
		} else {
			def["url"] = image.Url
			delete(def, "image")
		}
	} else {
		base64Str, err := image.Base64Encode()
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
