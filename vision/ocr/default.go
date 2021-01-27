package ocr

var defaultGeneralBasicParams = map[string]interface{}{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"url":              "",        //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效，不支持https的图片链接
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"probability":      "false",   //是否返回识别结果中每一行的置信度
}

var defaultGeneralWithLocationParams = map[string]interface{}{
	"image":                 "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"url":                   "",        //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效，不支持https的图片链接
	"recognize_granularity": "big",     //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"language_type":         "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction":      "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
	"detect_language":       "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"vertexes_location":     "false",   //是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
	"probability":           "false",   //是否返回识别结果中每一行的置信度
}

var defaultDeneralEnhancedParams = map[string]interface{}{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"probability":      "false",   //是否返回识别结果中每一行的置信度
}

var defaultWebimgParams = defaultDeneralEnhancedParams

var defaultIdcardParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"id_card_side":     "front", //front：身份证正面；back：身份证背面
	"detect_risk":      "false", //是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，默认不开启，即：false。可选值:true-开启；false-不开启
}

var defaultBankcardParams = map[string]interface{}{
	"image": "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
}

var defaultDriverLicenseParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
}

var defaultVehicleLicenseParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"accuracy":         "",      //normal 使用快速服务，1200ms左右时延；缺省或其它值使用高精度服务，1600ms左右时延
}

var defaultLicensePlateParams = defaultBankcardParams

var defaultFormParams = defaultBankcardParams

var defaultAccurateBasicParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"probability":      "false", //是否返回识别结果中每一行的置信度
}

var defaultVATInvoiceParams = map[string]interface{}{
	"image":    "",       //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"accuracy": "normal", //normal（默认配置）对应普通精度模型，识别速度较快，在四要素的准确率上和 high 模型保持一致，high对应高精度识别模型，相应的时延会增加，因为超时导致失败的情况也会增加（错误码282000）
	"type":     "normal", //进行识别的增值税发票类型，默认为 normal，可缺省, - normal：可识别增值税普票、专票、电子发票, - roll：可识别增值税卷票
}

var defaultIocrRecogniseParams = map[string]interface{}{
	"image":        "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"url":          "", //图片存储的 BOS（百度云存储）url，暂仅支持BOS url，不支持其他图床，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式，当image字段存在时url字段失效，不支持https的图片链接
	"templateSign": "", //模板 ID，自定义模板或预置模板的唯一标示，可用于调用指定的识别模板进行结构化识别，可在「模板管理」页查看并复制使用
	"classifierId": "", //分类器Id，分类器的唯一标示，可用于调用指定的分类器对传入的图片进行自动分类及识别，与 templateSign 至少存在一个，如同时存在，则优先级 templateSign > classfierId
}

var defaultIocrRecogniseFinanceParams = map[string]interface{}{
	"image":        "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"templateSign": "", //模板 ID，自定义模板或预置模板的唯一标示，可用于调用指定的识别模板进行结构化识别，可在「模板管理」页查看并复制使用
}

var defaultBusinessLicenseParams = map[string]interface{}{
	"image":            "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"url":              "", //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式，当image字段存在时url字段失效,请注意关闭URL防盗链
	"detect_direction": "", //此参数新版本无需传，支持自动检测图像旋转角度；朝向是指输入图像是正常方向、逆时针旋转90/180/270度
	"accuracy":         "", //可选值：normal,high参数选normal或为空使用快速服务；选择high使用高精度服务，但是时延会根据具体图片有相应的增加
}

var defaultCarTypeParams = map[string]interface{}{
	"image":     "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"top_num":   5,  //uint32	-	返回结果top n，默认5。
	"baike_num": 0,  //integer	0	返回百科信息的结果数，默认不返回
}

var defaultVinParams = map[string]interface{}{
	"image": "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"url":   "", //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式，当image字段存在时url字段失效,请注意关闭URL防盗链
}

var defaultNumberParams = map[string]interface{}{
	"image":                 "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	"recognize_granularity": "small", //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"detect_direction":      "false", //是否检测图像朝向，默认不检测，即：false。可选值包括true - 检测朝向；false - 不检测朝向。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。
}
