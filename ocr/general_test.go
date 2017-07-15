package ocr

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var apikey string = "nKesFdOqHYzKFMzVasdcvX4cfDU"
var secretkey string = "2xTn6TGucdffssUNa6YUZoDOMeZWqYsKpop1n"

var client *OCRClient = NewOCRClient(apikey, secretkey)

func TestGeneralRecognizeBasic(t *testing.T) {
	img, err := openfile("ocr.jpg")
	if err != nil {
		t.Fatal(err)
	}
	var conf map[string]string = make(map[string]string)
	bts, err := client.GeneralRecognizeBasic(img, conf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bts))
	t.Log("testing passed.")
}

func TestOCRClient_GeneralRecognizeWithLocation(t *testing.T) {
	img, err := openfile("ocr.jpg")
	if err != nil {
		t.Fatal(err)
	}
	var conf map[string]string = make(map[string]string)
	bts, err := client.GeneralRecognizeWithLocation(img, conf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bts))
	t.Log("testing passed.")
}

func TestOCRClient_GeneralRecognizeEnhanced(t *testing.T) {

	img, err := openfile("ocr.jpg")
	if err != nil {
		t.Fatal(err)
	}
	var conf map[string]string = make(map[string]string)
	bts, err := client.GeneralRecognizeEnhanced(img, conf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bts))
	t.Log("testing passed.")

}

func openfile(filename string) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return content, nil
}
