package study

import (
	"fmt"
	"testing"
)

// 列表增删改查
func TestList(t *testing.T) {
	//list
	list := make([]string, 0)
	list = append(list, "Hello", "World")
	fmt.Println(list[0])                 //获取
	fmt.Println(list[1:])                //截取
	list = append(list[:1], list[2:]...) //删除
	fmt.Println(list)
	var i = 1
	list = append(list[0:i], list[i-1:]...) //插入元素
	list[0] = "Ok"
	fmt.Println(list)
}
