package img

import (
	"github.com/jaryf/mutils/utils/xf"
	"strings"
)

func CheckHaveXrmn(x *xf.XfOcr, imgUrl, keyWord string) (res bool, err error) {
	wordList, err := x.ImgOcrXfFromUrl(imgUrl)
	if err != nil {
		return
	}
	for _, s := range wordList {
		contains := strings.Contains(s, keyWord)
		if contains {
			return true, nil
		}
	}
	return false, nil
}
