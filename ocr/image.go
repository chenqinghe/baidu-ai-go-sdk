package ocr

import (
	"image"
	"io"
)

type Size struct {
	Height int
	Width  int
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
