package study

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// 演示 chan 和 select 的使用
// 这里附带了一个超时机制
// chan 是线程安全的
func TestChan(t *testing.T) {
	index := make(chan int)
	go func() {
		time.Sleep(time.Second)
		index <- 1
	}()
	select {
	case v := <-index:
		log.Println(v)
	case <-time.After(5 * time.Second): // 不设置超时则会阻塞到 chan 存在数据
		fmt.Println("timeout") //超时
		return
	}
}
