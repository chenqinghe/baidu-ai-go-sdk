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
