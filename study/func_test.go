package study

import (
	"log"
	"testing"
)

// 我们可以在方法内修改上一级代码
func TestFunc(t *testing.T) {
	index := 1
	f := func() {
		index = 2
	}
	f()
	log.Println(index)
}
