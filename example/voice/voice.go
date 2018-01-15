package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"os"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

// Voice Composition
func TextToSpeech() {
	client := voice.NewVoiceClient(APIKEY, APISECRET)
	file, err := client.TextToSpeech("Hello World")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("hello.mp3", os.O_CREATE|os.O_WRONLY, 0644)
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

	f, err := os.OpenFile("16k.pcm", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	rs, err := client.SpeechToText(
		f,
		voice.Format("pcm"),
		voice.Channel(1),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
}

func main() {
	TextToSpeech()
	SpeechToText()
}
