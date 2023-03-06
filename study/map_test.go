package study

import (
	"fmt"
	"testing"
)

// map 增删改查
func TestMap(t *testing.T) {
	//map
	hashMap := make(map[string]string, 0)
	hashMap["A"] = "HelloWorld" //插入
	fmt.Println(hashMap["A"])   //获取
	delete(hashMap, "A")        //删除
	fmt.Println(hashMap)
}
