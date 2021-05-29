package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
	"sync"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "nKesFdOqHYzKFMzVcvX4cfDU"
	APISECRET = "2xTn6TGucUNa6YUZoDOMeZWqYsKpop1n"
)

var client *ocr.OCRClient

func init() {
	client = ocr.NewOCRClient(APIKEY, APISECRET)
}

func main() {
	//AccurateRecognizeBasic()
	//AccurateRecognizeBasic()
	//AccurateRecognize()
	GeneralRecognizeBasic()
	//HandWriting()
}

func GeneralRecognizeBasic() {
	rs, err := client.GeneralRecognizeBasic(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),       //是否检测图像朝向，默认不检测
		ocr.DetectLanguage(),        //是否检测语言，默认不检测。
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
		ocr.WithProbability(),       //是否返回识别结果中每一行的置信度
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}

func HandWriting() {
	rs, err := client.HandWriting(
		vision.MustFromFile("ocr_hand_writing.jpg"),
		ocr.WithProbability(), //是否返回识别结果中每一行的置信度
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}

func ConcurrentAccurateRecognizeBasic() {
	images := []string{
		"ocr_image_0.jpg",
		"ocr_image_1.jpg",
		"ocr_image_2.jpg",
		"ocr_image_3.jpg",
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(images))
	for _, image := range images {
		go func(img string) {
			defer wg.Done()
			resp, err := client.AccurateRecognizeBasic(
				vision.MustFromFile(img),
				ocr.DetectDirection(),
				ocr.WithProbability(),
			)
			if err != nil {
				panic(err)
			}
			fmt.Println(img)
			fmt.Println(resp.ToString())
		}(image)
	}
	wg.Wait()
}

func AccurateRecognizeBasic() {
	resp, err := client.AccurateRecognizeBasic(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),
		ocr.WithProbability(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ToString())
}

func AccurateRecognize() {
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
