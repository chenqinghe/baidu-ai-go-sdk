package image

import "context"

func (e *ImageClient) Colourize(ctx context.Context, param *Input) (*EnhanceResponse, error) {
	return e.posUrlEncode(ctx, urlColorize, param, image8M)
}
