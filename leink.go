package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	next *Student
}

// Add 编写第一种插入方法，在单链表的最后加入
func (stu *Student) Add(new *Student) {
	temp := stu
	for {
		if temp.next == nil {
			break //已找到最后
		}
		temp = temp.next //
	}
	temp.next = new
}

//按从小到大添加
func (stu *Student) Xadd(new *Student) {
	temp := stu
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.id > new.id {
			break
		} else if temp.next.id == new.id {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag {
		new.next = temp.next
		temp.next = new
	} else {
		fmt.Println("有相同的数据")
	}
}

//删除一个元素
func (stu *Student) Delect(id int) {
	temp := stu
	flag := false
	for {
		if temp.next == nil {
			flag = false
			break
		} else if temp.next.id == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("没有数据")
	}
}

//显示
func (stu *Student) Show() {
	temp := stu
	if temp.next == nil {
		fmt.Println("什么都没有")
		return
	}
	for {
		fmt.Printf("id=%d,名字=%s\n", temp.next.id, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}

	}
}

func main() {
	stu := &Student{}
	stu1 := &Student{
		id:   1,
		name: "宋江",
	}
	stu2 := &Student{
		id:   2,
		name: "卢俊义",
	}
	stu3 := &Student{
		id:   3,
		name: "张飞",
	}

	stu.Xadd(stu2)
	stu.Xadd(stu1)
	stu.Xadd(stu3)

	stu.Show()

}
