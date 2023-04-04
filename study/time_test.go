package study

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// 单次定时器
	timer := time.NewTimer(time.Second * 3)
	stop := time.NewTimer(time.Second * 3 * 2)
	for {
		select {
		case <-timer.C:
			t.Log("timer")
			// 重置定时器为1秒
			timer.Reset(time.Second * 1)
		case <-stop.C:
			t.Log("stop")
			return
		}

	}

}

func TestUpdate(t *testing.T) {

	fmt.Println(time.Second / 20)

	// 每秒20帧
	timer := time.NewTicker(time.Second)
	// 重置定时器
	for {
		select {
		case <-timer.C:
			t.Log("timer")
			// 重置定时器为1秒
			timer.Reset(time.Second)
		}
	}

}
