package voice

import (
	"testing"
)

var (
	apikey    = "MDNsII2jkUtbF729GQOZt7FS"
	secretkey = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

var client *VoiceClient

func init() {
	client = NewVoiceClient(apikey, secretkey)
}

func TestVoiceClient_TextToSpeech(t *testing.T) {
	res, err := client.TextToSpeech(
		"你好",
		Speed(10),
		Person(1),
		Volume(10),
	)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	t.Log(res)
}
