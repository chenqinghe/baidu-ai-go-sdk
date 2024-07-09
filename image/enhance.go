package image

import (
	"context"
)

// ContrastEnhance https://ai.baidu.com/ai-doc/IMAGEPROCESS/ek3bclnzn
func (e *ImageClient) ContrastEnhance(ctx context.Context, param *Input) (*EnhanceResponse, error) {
	return e.posUrlEncode(ctx, urlContrastEnhance, param, image8M)
}

func (e *ImageClient) ColorEnhance(ctx context.Context, param *Input) (*EnhanceResponse, error) {
	return e.posUrlEncode(ctx, urlColorEnhance, param, image8M)
}

func (e *ImageClient) QualityEnhance(ctx context.Context, param *Input) (*EnhanceResponse, error) {
	return e.posUrlEncode(ctx, urlQualityEnhance, param, image4M)
}

func (e *ImageClient) DefinitionEnhance(ctx context.Context, param *Input) (*EnhanceResponse, error) {
	return e.posUrlEncode(ctx, urlDefinitionEnhance, param, image8M)
}
