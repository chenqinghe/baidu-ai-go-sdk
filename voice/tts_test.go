package voice

import (
	"testing"
)

var (
	apikey    = "MDNsII2jkUtbF729GQOZt7FS"
	secretkey = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

var client *VoiceClient

func TestNewVoiceClient(t *testing.T) {
	client = NewVoiceClient(apikey, secretkey)
}

func TestVoiceClient_TextToSpeech(t *testing.T) {
	_, err := client.TextToSpeech("你好")
	if err != nil {
		t.Fail()
	}
}

func TestVoiceClient_TextToSpeech2(t *testing.T) {
	_,err:=client.SpeechToText()
}
