package main

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"os"
	"io/ioutil"
	"encoding/base64"
	"fmt"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
	APISECRET = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"
)

// Voice Composition
func TextToSpeech() {
	client := voice.NewVoiceClient(APIKEY, APISECRET)
	file, err := client.TextToSpeech("Hello World")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("hello.mp3", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(file); err != nil {
		log.Fatal(err)
	}
}

// Voice Recognition
// ATTENTION: the .wav file must be 8k or 16k rate with single(mono) channel.
// FYI: you can use QuickTime to record voice and Fission converting to .wav 
func SpeechToText() {
	client := voice.NewVoiceClient(APIKEY, APISECRET)
	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("hello.wav", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fi, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		log.Fatal(err1)
	}
	afterBase64Str := base64.StdEncoding.EncodeToString(fi)
	fiLen := len(fi)
	param := voice.ASRParams{
		Format: "wav",
		Rate: 16000,
		Channel: 1,
		Cuid: "12312312112", 
		Token: client.AccessToken,
		Lan: "zh",
		Speech: afterBase64Str,
		Len: fiLen,
	}
	rs, err2 := client.SpeechToText(param)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(rs)
}

func main() {
	TextToSpeech()
	SpeechToText()
}
