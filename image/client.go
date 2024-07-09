package image

import (
	"context"
	"fmt"
	gosdk "github.com/chenqinghe/baidu-ai-go-sdk"
	"github.com/imroc/req"
	"net/http"
)

type ImageClient struct {
	*gosdk.Client
}

func NewEnhanceClient(apiKey, secretKey string) *ImageClient {
	return &ImageClient{
		Client: gosdk.NewClient(apiKey, secretKey),
	}
}

func (e *ImageClient) requestUrl(url string) (string, error) {
	if err := e.Auth(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s?access_token=%s", url, e.AccessToken), nil
}

func (e *ImageClient) posUrlEncode(ctx context.Context, url string,
	inputParam *Input, maxSize int) (*EnhanceResponse, error) {
	options, err := inputParam.encode(maxSize)
	if err != nil {
		return nil, err
	}

	url, err = e.requestUrl(url)
	if err != nil {
		return nil, err
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.
		Post(url, req.Param(options), header, ctx)
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != http.StatusOK {
		return nil, APIError{
			ErrorCode:  0,
			ErrorMsg:   resp.String(),
			StatusCode: resp.Response().Status,
		}
	}

	var response generalResponse
	err = resp.ToJSON(&response)
	if err != nil {
		return nil, fmt.Errorf("decode enhance response: %w", err)
	}

	if response.success() {
		return &response.EnhanceResponse, nil
	}

	return nil, response.APIError
}
