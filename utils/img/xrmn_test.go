package img

import (
	"github.com/jaryf/mutils/utils/xf"
	"testing"
)

const (
	APPID     = "5dedf6f5"
	APISecret = "82a5c5cde0407a20204a15f017687062"
	APIKey    = "2a44bb0136b097b083042ea86882002e"
	xfWebAPI  = "https://api.xf-yun.com/v1/private/sf8e6aca1"
	xfHost    = "api.xf-yun.com"
	imgUrl    = "https://t.xrmn5.cc/UploadFile/202204/22/5D155510249.jpg"
	imgUrl2   = "https://t.xrmn5.cc/UploadFile/202204/22/A3155511156.jpg"
	imgPath   = "5D155510249.jpg"
)

var xfClient *xf.XfOcr

func init() {
	xfClient = xf.NewXfOcr(APPID, APISecret, APIKey, xfWebAPI, xfHost)
}

func TestCheckHaveXrmn(t *testing.T) {
	res, err := CheckHaveXrmn(xfClient, imgUrl2, "xrmn")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
