package image

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func decodeToLocal(response *EnhanceResponse) error {
	binary, err := base64.StdEncoding.DecodeString(response.Image)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("./test_result/%d.png", time.Now().UnixMilli()), binary, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return nil
}
