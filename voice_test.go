package dueros

import (
	"fmt"
	"os"
	"testing"
)

var apiKey string = "tVPKdPqxKwWOM9vsukPzseoH"
var apiSecret string = "VFusINaQ3YjDUC4GoIB1cENME9g2f4Gn"

func TestDefaultAuthorizer_Authorize(t *testing.T) {
	client := NewVoiceClient(apiKey, apiSecret)
	if err := client.auth(); err != nil {
		t.Fatal(err)
	}
	if client.AccessToken == "" {
		t.Error("获取access_token失败")
	}
	fmt.Println(client.AccessToken)
	t.Log("testing passed.")
}

func TestVoiceClient_TextToSpeech(t *testing.T) {
	client := NewVoiceClient(apiKey, apiSecret)
	file, err := client.TextToSpeech("哈哈哈哈哈哈哈哈哈")
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.OpenFile("hello.mp3", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(file); err != nil {
		t.Fatal(err)
	}
	t.Log("testing passed.")
}
