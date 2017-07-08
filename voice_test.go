package dueros

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var apiKey string = "tVPKdPqxKwWOasdM9vsukPzseoHhk"
var apiSecret string = "VFusINaQ3YjDUC4GoIB1casdENME9g2f4Gn"

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

func TestDefaultAuthorizer_Authorize_fail(t *testing.T) {
	client := NewVoiceClient("", apiSecret)
	if err := client.auth(); err != nil {
		t.Log(err)
	}
	t.Log("testing passed.")
}

func TestVoiceClient_TextToSpeech(t *testing.T) {
	client := NewVoiceClient(apiKey, apiSecret)
	file, err := client.UseDefaultTTSConfig().TextToSpeech("你叫什么名字啊？")
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

func TestVoiceClient_SpeechToText(t *testing.T) {
	client := NewVoiceClient(apiKey, apiSecret)
	f, err := os.OpenFile("hello.wav", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	fi, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	afterBase64Str := base64.StdEncoding.EncodeToString(fi)
	fiLen := len(fi)
	param := ASRParams{
		Format:  "wav",
		Rate:    16000,
		Channel: 1,
		Cuid:    "12312312112",
		Token:   client.AccessToken,
		Lan:     "zh",
		Speech:  afterBase64Str,
		Len:     fiLen,
	}
	fmt.Println(param.Token)
	rs, err := client.SpeechToText(param)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rs)
	t.Log("testing passed.")
}
