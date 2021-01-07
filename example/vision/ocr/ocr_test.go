package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
	"testing"
)

func init() {
	client = ocr.NewOCRClient(APIKEY, APISECRET)
}

func TestAccurateRecognize(t *testing.T) {
	resp, err := client.AccurateRecognize(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),
		ocr.WithProbability(),
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ToString())
}

func TestAccurateRecognizeBasic(t *testing.T) {
	resp, err := client.AccurateRecognizeBasic(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),
		ocr.WithProbability(),
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ToString())
}
