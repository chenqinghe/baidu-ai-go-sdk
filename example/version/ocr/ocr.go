package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/version/ocr"
	"os"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
	APISECRET = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"
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
