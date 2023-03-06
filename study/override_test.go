package study

import (
	"log"
	"testing"
)

type Cat interface {
	PrintInfo()
}
type BlueCat struct {
}

func (b BlueCat) PrintInfo() {
	log.Println("蓝色猫")
}

func (b BlueCat) Color() {
	log.Println("蓝色")
}

type OrangeCat struct {
	BlueCat
}

func (b OrangeCat) PrintInfo() {
	log.Println("橙色猫")
}

// 覆盖测试
func TestOverride(t *testing.T) {
	var o = &OrangeCat{}
	o.PrintInfo() // 重写的方法
	o.Color()     // 继承的方法
}
