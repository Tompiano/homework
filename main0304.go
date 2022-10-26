package main

import "fmt"

type Interface1 interface {
	BubbleSort()
}

type BubbleSort struct {
	str []int
}

func (s *BubbleSort) Print1(str []int) {
	//冒泡排序
	for i := 0; i < len(str); i++ {

		for j := 0; j < len(str)-i-1; j++ {
			if str[j] > str[j+1] {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}
	fmt.Println("排序后", str)

}

func main() {
	str := []int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	var s BubbleSort
	s.Print1(str) //使用方法

}
