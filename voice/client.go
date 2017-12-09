package voice

import "github.com/chenqinghe/baidu-ai-go-sdk"

const (
	B = 1 << 10 * iota
	KB
	MB
	GB
	TB
	PB
)

type VoiceClient struct {
	*gosdk.Client
}

func NewVoiceClient(apiKey, apiSecret string) *VoiceClient {
	return &VoiceClient{
		Client: gosdk.NewClient(apiKey, apiSecret),
	}
}
