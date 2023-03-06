package study

import (
	"log"
	"testing"
)

// Go 多线程
// 指在于多线程任务时使用,例如调用远程 A,B,C 接口各耗时1s 串行需要3s
// 例如 chan 就可以启用3个线程并同时等待他们的结果,直至请求完成。那么只需要1s多,比串行效率大大提高
func TestGo(t *testing.T) {
	index := make(chan int)
	for i := 0; i < 2; i++ {
		v := i
		go func() {
			index <- v
		}()
	}
	for i := 0; i < 2; i++ {
		select {
		case v := <-index:
			log.Println(v)
		}
	}
}
