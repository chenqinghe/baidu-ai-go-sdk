package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/version/ocr"
	"os"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

func main() {

	client := ocr.NewOCRClient(APIKEY, APISECRET)

	f, err := os.OpenFile("ocr.jpg", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	rs, err := client.GeneralRecognizeBasic(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rs))
}
