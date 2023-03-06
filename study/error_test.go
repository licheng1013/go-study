package study

import (
	"log"
	"testing"
)

// 演示如何捕获异常
func TestError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%v", r)
		}
	}()
	panic("异常示例1")
}
