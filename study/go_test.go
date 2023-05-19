package study

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
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

// 测试一组线程使用
func TestGroupGO(t *testing.T) {
	// sync.WaitGroup 用于需要等待一组任务完成的场景
	// 如根据用户id,查询用户信息,查询用户订单.采用多线程效率更快
	var v sync.WaitGroup
	for i := 0; i < 3; i++ {
		v.Add(1)
		j := i
		go func() {
			time.Sleep(time.Duration(j) * time.Second)
			log.Println("go", j)
			v.Done()
		}()
	}
	v.Wait()
	fmt.Println("等待执行完毕")
}
