package goFunTest

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestIsNo(t *testing.T) {
	val := 18.0
	res := IsNo(val)
	spew.Dump(res)
}