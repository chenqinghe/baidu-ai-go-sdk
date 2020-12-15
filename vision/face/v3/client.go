package v3

import (
	sdk "github.com/chenqinghe/baidu-ai-go-sdk"
)

type FaceClient struct {
	*sdk.Client
}

func NewFaceClient(key, secret string) *FaceClient {
	return &FaceClient{
		Client: sdk.NewClient(key, secret),
	}
}
