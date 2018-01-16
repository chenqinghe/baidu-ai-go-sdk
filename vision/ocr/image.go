package ocr

import (
	"bytes"
	"encoding/base64"

	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

const (
	MAX_SIZE = 4096
	MIN_SIZE = 15
)

type Image struct {
	Reader io.Reader
	Url    string
	Size   *Size
}

type Size struct {
	Height int
	Width  int
}

func FromFile(file string) (*Image, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	imageContent, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	size, err := getImageSize(bytes.NewBuffer(imageContent))
	if err != nil {
		return nil, err
	}
	if size.Height > MAX_SIZE || size.Height < MIN_SIZE || size.Width > MAX_SIZE || size.Width < MIN_SIZE {
		return nil, errors.New("image size is invalid")
	}

	reader := bytes.NewBuffer(imageContent)

	return &Image{
		Reader: reader,
		Size:   size,
	}, nil

}

func MustFromFile(file string) *Image {
	img, err := FromFile(file)
	if err != nil {
		panic(err)
	}
	return img
}

func FromReader(reader io.Reader) (*Image, error) {

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	size, err := getImageSize(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	if size.Height > MAX_SIZE || size.Height < MIN_SIZE || size.Width > MAX_SIZE || size.Width < MIN_SIZE {
		return nil, errors.New("image size is invalid")
	}

	rd := bytes.NewReader(content)

	return &Image{
		Reader: rd,
		Size:   size,
	}, nil
}

func MustFromReader(reader io.Reader) *Image {
	img, err := FromReader(reader)
	if err != nil {
		panic(err)
	}
	return img
}

func FromBytes(bts []byte) (*Image, error) {
	buf := bytes.NewReader(bts)
	size, err := getImageSize(buf)
	if err != nil {
		return nil, err
	}
	if size.Height > MAX_SIZE || size.Height < MIN_SIZE || size.Width > MAX_SIZE || size.Width < MIN_SIZE {
		return nil, errors.New("image size is invalid")
	}

	reader := bytes.NewReader(bts)

	return &Image{
		Reader: reader,
		Size:   size,
		Url:    "",
	}, nil

}

func MustFromBytes(bts []byte) *Image {
	img, err := FromBytes(bts)
	if err != nil {
		panic(err)
	}
	return img
}

func FromUrl(link string) (*Image, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "https" {
		return nil, errors.New("not support https scheme.")
	}
	return &Image{
		Reader: nil,
		Url:    link,
		Size:   nil,
	}, nil
}

func MustFromUrl(link string) *Image {
	img, err := FromUrl(link)
	if err != nil {
		panic(err)
	}
	return img
}

func (img *Image) Base64() (string, error) {

	if img.Reader == nil {
		if img.Url != "" {
			return "", errors.New("can not encode image specified by url")
		} else {
			return "", errors.New("no image source detected")
		}
	}

	bts, err := ioutil.ReadAll(img.Reader)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bts), nil

}

func getImageSize(reader io.Reader) (*Size, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	size := &Size{
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
	}
	return size, nil
}
