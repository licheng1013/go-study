package main

import "testing"

func Test1(t *testing.T) {
	strA := "hello"
	length := len(strA)
	strB := ""
	for index := range strA {
		strB += strA[length-index-1 : length-index]
	}

	// 验证
	for i := range strA {
		strAChar := strA[length-i-1 : length-i]
		strBChar := strB[i : i+1]
		if strAChar != strBChar {
			t.Error("反转字符串失败")
		}
	}
	t.Log(strA, "\t", strB)
}
