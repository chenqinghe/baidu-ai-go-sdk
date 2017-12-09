package ocr

const (
	OCR_WEBIMAGE_URL       = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"
	OCR_IDCARD_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	OCR_BANKCARD_URL       = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	OCR_DRIVERLICENSE_URL  = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
	OCR_VEHICLELICENSE_URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"
	OCR_LICENSEPLATE_URL   = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	OCR_FORM_URL           = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"
)

var (
	defaultWebimgParams = defaultDeneralEnhancedParams
	defaultIdcardParams = map[string]string{
		"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
		"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
		"id_card_side":     "front", //front：身份证正面；back：身份证背面
		"detect_risk":      "false", //是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，默认不开启，即：false。可选值:true-开启；false-不开启
	}
	defaultBankcardParams = map[string]string{
		"image": "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	}
	defaultDriverLicenseParams = map[string]string{
		"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
		"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	}
	defaultVehicleLicenseParams = map[string]string{
		"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
		"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
		"accuracy":         "",      //normal 使用快速服务，1200ms左右时延；缺省或其它值使用高精度服务，1600ms左右时延
	}
	defaultLicensePlateParams = defaultBankcardParams
	defaultFormParams         = defaultBankcardParams
)

func (oc *OCRClient) WebImageRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_WEBIMAGE_URL, conf, defaultWebimgParams)

}

func (oc *OCRClient) IdcardRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_IDCARD_URL, conf, defaultIdcardParams)
}

func (oc *OCRClient) BankcardRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_BANKCARD_URL, conf, defaultBankcardParams)

}

func (oc *OCRClient) DriverLicenseRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_DRIVERLICENSE_URL, conf, defaultDriverLicenseParams)
}

func (oc *OCRClient) VehicleLicenseRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_VEHICLELICENSE_URL, conf, defaultVehicleLicenseParams)
}

func (oc *OCRClient) LicensePlateRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_LICENSEPLATE_URL, conf, defaultLicensePlateParams)
}

func (oc *OCRClient) FromdataRecognize(img []byte, conf map[string]string) ([]byte, error) {

	return oc.generalOperate(img, OCR_FORM_URL, conf, defaultFormParams)
}

func (oc *OCRClient) generalOperate(img []byte, baseurl string, conf, def map[string]string) ([]byte, error) {
	if err := oc.Auth(); err != nil {
		return nil, err
	}
	conf = parseParams(def, conf)

	url := baseurl + "?access_token=" + oc.AccessToken

	return doRequest(url, conf)
}
