package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/face"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

func main() {
	client := face.NewFaceClient(APIKEY, APISECRET)

	rs, err := client.Match(
		vision.MustFromFile("p1.jpg"),
		vision.MustFromFile("p2.jpg"),
		map[string]interface{}{
			"ext_fields":     "qualities",                 //返回质量信息，取值固定，目前支持qualities(质量检测)(对所有图片都会做改处理)
			"image_liveness": "faceliveness,faceliveness", //返回的活体信息，“faceliveness,faceliveness” 表示对比对的两张图片都做活体检测；“,faceliveness” 表示对第一张图片不做活体检测、第二张图做活体检测；“faceliveness,” 表示对第一张图片做活体检测、第二张图不做活体检测；
			"types":          "7,7",
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(rs.Dump())
}
