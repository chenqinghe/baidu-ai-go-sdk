package main

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"os"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
	APISECRET = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"
)

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

func main() {
	TextToSpeech()
}
