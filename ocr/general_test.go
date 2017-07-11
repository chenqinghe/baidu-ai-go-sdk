package ocr

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var apikey string = "nKesFsdfgdOqHYzKFMzVcvdfX4cfDU"
var secretkey string = "2xTn6TGucUNa6YfgfdUZoDOMeZWqYsKpop1n"

func TestGeneralRecognizeBasic(t *testing.T) {
	f, err := os.OpenFile("ocr.jpg", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	img, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	client := NewOCRClient(apikey, secretkey)
	var conf map[string]string = make(map[string]string)
	bts, err := client.GeneralRecognizeBasic(img, conf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bts))
	t.Log("testing passed.")
}
