package study

import (
	"fmt"
	"sort"
	"testing"
)

// 排序
type StudentTasks []StudentTask

func (s StudentTasks) Len() int {
	return len(s)
}
func (s StudentTasks) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}
func (s StudentTasks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type StudentTask struct {
	Score int
}

// map 增删改查
func TestSort(t *testing.T) {
	// 创建一个学生切片
	students := StudentTasks{
		{90},
		{60},
		{80},
		{76},
	}
	sort.Sort(students)
	fmt.Println(students) // [{90} {80} {76} {60}]
}
