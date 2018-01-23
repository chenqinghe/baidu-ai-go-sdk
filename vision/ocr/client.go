package ocr

import "github.com/chenqinghe/baidu-ai-go-sdk"

type OCRClient struct {
	*gosdk.Client
}

func NewOCRClient(apiKey, secretKey string) *OCRClient {
	return &OCRClient{
		Client: gosdk.NewClient(apiKey, secretKey),
	}
}
