package goFunTest

import (
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
	"testing"
)

func TestParseUnicode(t *testing.T){
	str := "md5\\u9519\\u8bef"

	str = "\\u64cd\\u4f5c\\u8fdd\\u89c4,\\u95f4\\u96943\\u79d2\\u540e\\u518d\\u8bf7\\u6c42"

	strBytes, err := UnescapeUnicode([]byte(str))

	spew.Dump(string(strBytes), err)

}


func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

