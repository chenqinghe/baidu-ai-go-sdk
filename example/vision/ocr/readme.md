# Instanlation 
```Go
go get github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr
```
# Usage
- 创建client

```Go
const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

client := ocr.NewOCRClient(APIKEY, APISECRET)

```

- 通用文字识别
```Go
f, err := os.OpenFile("ocr.jpg", os.O_RDONLY, 0777)
if err != nil {
  panic(err)
}
rs, err := client.GeneralRecognizeBasic(
  f,//必须
  LanguageType("CHN_ENG"),//非必须，语言类型，默认CHN_ENG
  DetectDirection(),//非必须，检测图像朝向，默认不检测
  DetectLanguage(),//非必须，检测语言，默认不检测
  WithProbability(),//非必须，返回识别结果中每一行的置信度
)

```
