package image

import "github.com/chenqinghe/baidu-ai-go-sdk/voice"

const (
	urlContrastEnhance   = "https://aip.baidubce.com/rest/2.0/image-process/v1/contrast_enhance"
	urlColorEnhance      = "https://aip.baidubce.com/rest/2.0/image-process/v1/color_enhance"
	urlColorize          = "https://aip.baidubce.com/rest/2.0/image-process/v1/colourize"
	urlQualityEnhance    = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_quality_enhance"
	urlDefinitionEnhance = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_definition_enhance"
)

const (
	image8M = 8 * voice.MB
	image4M = 4 * voice.MB
)
