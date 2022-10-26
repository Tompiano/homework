package main

import (
	"fmt"
	"strings"
)

type Separate interface {  //建立接口
	Stats()
	Find()
}

type Coder struct {
	Name, birth      string
	Id, Number, Code int
}

func (a Coder) Stats() (Number int, stu []string) {
	//实现简单的统计功能，比如统计做题数大于150的同学并输出其完整信息
	for i := 0; i < 3; i++ {
		if Number > 150 {
			fmt.Println(stu[i])
		}
	}
	return
}

func (a Coder) Find() (stu []string) {
	//按姓名查找：如找Jay
	for i := 0; i < len(stu); i++ {
		index1 := strings.Index(stu[i], "Jay")
		if index1 != -1 {
			fmt.Println(stu[i])
		}
	}
	//按学号查找：如找id=1的同学
	for i := 0; i < len(stu); i++ {
		index2 := strings.Index(stu[i], "1")
		if index2 != -1 {
			fmt.Println(stu[i])
		}
	}
	return
}

func main() {
	var stu [3]Coder = [3]Coder{

		//实现对多个同学的信息输入，输出
		Coder{"Jay", "7月1日", 1, 234, 100},
		Coder{"John", "8月1日", 2, 23, 200},
		Coder{"Jane", "9月1日", 3, 456, 300},
	}
	//实现信息修改功能，如：修改名字
	stu[1].Name = "大灰狼"

	fmt.Println(stu)
	var a Coder
	a.Stats()
	a.Find()

}
