package leetcode

import (
	"strings"
	"testing"
)

// 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，
//判断字符串是否有效。有效字符串需满足：左括号必须用相同类型的右括号闭合，
//左括号必须以正确的顺序闭合。
//例如，输入字符串 "()[]{}"，输出 true；输入字符串 "([)]"，输出 false

func Test4(t *testing.T) {
	// 定义括号
	strList := "()[]{}"
	str := "(xxx)aa{aa}{aa}aa" // TODO 存在问题

	current := make([]string, 0)
	for i, o := range str {
		k := string(o)
		//t.Logf("字符串: %v", k)
		index := strings.Index(strList, k)
		if i == 0 && index == -1 {
			t.Log("非法字符串: ", k)
		}
		if index != -1 {
			if index == 0 || index == 2 || index == 4 {
				// 左边括号
				current = append(current, k)
			} else {
				// 右边括号
				leftStrIndex := strings.Index(strList, current[len(current)-1])
				// 当前是右阔后则
				if leftStrIndex == 0 && index == 1 {
					current = current[:len(current)-1]
				} else if leftStrIndex == 2 && index == 3 {
					current = current[:len(current)-1]
				} else if leftStrIndex == 4 && index == 5 {
					current = current[:len(current)-1]
				} else {
					t.Logf("闭合错误:  %s-%s", current, k)
					return
				}
			}
		}

	}
}
