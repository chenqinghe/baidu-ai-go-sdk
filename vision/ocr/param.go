package ocr

type RequestParam func(map[string]interface{})

//识别语言类型，默认为CHN_ENG。
func LanguageType(lang string) RequestParam {
	options := []string{
		"CHN_ENG",
		"ENG",
		"POR",
		"FRE",
		"GER",
		"ITA",
		"SPA",
		"RUS",
		"JAP",
		"KOR",
	}

	illegal := true
	for _, v := range options {
		if v == lang {
			illegal = false
			break
		}
	}

	if illegal {
		lang = "CHN_ENG"
	}
	return func(m map[string]interface{}) {
		m["language_type"] = lang
	}
}

//是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
//- true：检测朝向；
//- false：不检测朝向。
func DetectDirection() RequestParam {
	return func(m map[string]interface{}) {
		m["detect_direction"] = true
	}
}

//是否检测语言，默认不检测。
//当前支持（中文、英语、日语、韩语）
func DetectLanguage() RequestParam {
	return func(m map[string]interface{}) {
		m["detect_language"] = true
	}
}

//是否返回识别结果中每一行的置信度
func WithProbability() RequestParam {
	return func(m map[string]interface{}) {
		m["probability"] = true
	}
}

//是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
func RecognizeGranularity() RequestParam {
	return func(m map[string]interface{}) {
		m["recognize_granularity"] = "small"
	}
}

//是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
func WithVertexesLocation() RequestParam {
	return func(m map[string]interface{}) {
		m["vertexes_location"] = true
	}
}

//front：身份证含照片的一面；back：身份证带国徽的一面
func IDCardSide(side string) RequestParam {
	return func(m map[string]interface{}) {
		m["id_card_side"] = side
	}
}

//是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，默认不开启，即：false。
// 可选值:true-开启；false-不开启
func DetectRisk() RequestParam {
	return func(m map[string]interface{}) {
		m["detect_risk"] = true
	}
}

//true: 归一化格式输出；false 或无此参数按非归一化格式输出
func UnifiedValidPeriod() RequestParam {
	return func(m map[string]interface{}) {
		m["unified_valid_period"] = true
	}
}

//normal 使用快速服务，1200ms左右时延；缺省或其它值使用高精度服务，1600ms左右时延
func Accuracy(opt string) RequestParam {
	if opt != "normal" && opt != "high" {
		opt = "normal"
	}
	return func(m map[string]interface{}) {
		m["accuracy"] = opt
	}
}

//是否检测多张车牌，默认为false，当置为true的时候可以对一张图片内的多张车牌进行识别
func MultiDetect() RequestParam {
	return func(m map[string]interface{}) {
		m["multi_detect"] = true
	}
}

//自定义模板文字识别 模板号
func TemplateSign(templateSign string) RequestParam {
	return func(m map[string]interface{}) {
		m["templateSign"] = templateSign
	}
}

//自定义模板文字识别 分类器Id
func ClassifierId(classifierId int) RequestParam {
	return func(m map[string]interface{}) {
		m["classifierId"] = classifierId
	}
}
