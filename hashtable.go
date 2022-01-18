package main

import (
	"fmt"
	"os"
)

//定义员工结构体
type Emp struct {
	id   int
	name string
	Next *Emp
}

//添加数据
func (em *Emp) Insert(emp *Emp) {
	mp := em
	for {
		if mp.Next == nil {
			break
		}
		if mp.id > emp.id {
			break
		}
		mp = mp.Next
	}
	mp.Next = emp
}

//根据id查找对应的雇员，如果没有就返回nil
func (emp *Emp) FindById(id int) *Emp {
	cur := emp
	for {
		if cur != nil && cur.id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}
func (emp *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d\n", emp.id%7, emp.id)
}

func (em *Emp) show() {
	mp := em
	if mp.Next == nil {
		fmt.Println("什么都没有")
		return
	}
	for {
		fmt.Printf("id=%d,名字=%s\n", mp.Next.id, mp.Next.name)
		mp = mp.Next
		if mp.Next == nil {
			break
		}
	}

}

//定义含有一个链表的数组
type HashTable struct {
	LinkArr [7]Emp
}

//编写一个散列方法
func (hash *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对应的链表的下标
}

//给HashTable 编写Insert 雇员的方法.
func (hash *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := hash.HashFun(emp.id)
	//使用对应的链表添加
	hash.LinkArr[linkNo].Insert(emp) //
}

//编写方法，显示hashtable的所有雇员
func (hash *HashTable) ShowAll() {
	for i := 0; i < len(hash.LinkArr); i++ {
		hash.LinkArr[i].ShowLink(i)
	}
}

//显示链表的信息
func (emp *Emp) ShowLink(no int) {
	if emp.Next == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	//变量当前的链表，并显示数据
	cur := emp // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.id, cur.name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理

}

//编写一个方法，完成查找
func (hash *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定将该雇员应该在哪个链表
	linkNo := hash.HashFun(id)
	return hash.LinkArr[linkNo].FindById(id)
}
func main() {

	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("1 表示添加雇员")
		fmt.Println("2  表示显示雇员")
		fmt.Println("3  表示查找雇员")
		fmt.Println("4  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				id:   id,
				name: name,
			}
			hashtable.Insert(emp)
		case "2":
			hashtable.ShowAll()
		case "3":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}

		case "4":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}

}
