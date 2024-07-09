package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var (
	client *ImageClient
)

func preTest() {
	if client != nil {
		return
	}
	client = NewEnhanceClient(os.Getenv("BAIDU_API_KEY"),
		os.Getenv("BAIDU_SECRET_KEY"))
	err := os.MkdirAll("./test_result", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestAllAPI(t *testing.T) {
	preTest()

	t.Run("contrastEnhance", TestContrastEnhance)

	t.Run("color_enhance", func(t *testing.T) {
		err := generalCall(client.ColorEnhance, &Input{
			ImageUrl:    "https://ai.bdstatic.com/file/75F5ABC751594F55B23AC4168F4A919A",
			File:        nil,
			ImageBase64: "",
		})
		require.NoError(t, err)
	})

	t.Run("quality_enhance", func(t *testing.T) {
		err := generalCall(client.QualityEnhance, &Input{
			ImageUrl:    "https://ai-public-console.cdn.bcebos.com/portal-pc-static/1720085025198/images/technology/imageprocess/image_quality_enhance/1.jpg",
			File:        nil,
			ImageBase64: "",
		})
		require.NoError(t, err)
	})

	t.Run("definition_enhance", func(t *testing.T) {
		err := generalCall(client.DefinitionEnhance, &Input{
			ImageUrl:    "https://ai-public-console.cdn.bcebos.com/portal-pc-static/1720085025198/images/technology/imageprocess/image_definition_enhance/1-1.jpg",
			File:        nil,
			ImageBase64: "",
		})
		require.NoError(t, err)
	})

	t.Run("colourize", func(t *testing.T) {
		err := generalCall(client.Colourize, &Input{
			ImageUrl:    "https://ai-public-console.cdn.bcebos.com/portal-pc-static/1720085025198/images/technology/imageprocess/colourize/1.jpg",
			File:        nil,
			ImageBase64: "",
		})
		require.NoError(t, err)
	})
}

func TestContrastEnhance(t *testing.T) {
	preTest()

	err := generalCall(client.ContrastEnhance, &Input{
		ImageUrl:    "https://ai-public-console.cdn.bcebos.com/portal-pc-static/1720085025198/images/technology/imageprocess/contrast_enhance/1.jpg",
		File:        nil,
		ImageBase64: "",
	})
	require.NoError(t, err)

	binary, err := ioutil.ReadFile("../example/image/enhance_01.jpg")
	require.NoError(t, err)

	err = generalCall(client.ContrastEnhance, &Input{
		ImageBase64: base64.StdEncoding.EncodeToString(binary),
	})
	require.NoError(t, err)

	err = generalCall(client.ContrastEnhance, &Input{
		File: bytes.NewBuffer(binary),
	})
	require.NoError(t, err)

	err = generalCall(client.ContrastEnhance, &Input{
		ImageUrl: "https://www.baidu.com",
	})
	require.NotNil(t, err, err)
}

func generalCall(fn func(ctx context.Context, input *Input) (*EnhanceResponse, error), input *Input) error {
	resp, err := fn(context.TODO(), input)
	if err != nil {
		return err
	}
	err = decodeToLocal(resp)
	if err != nil {
		return err
	}
	return nil
}
