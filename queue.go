package main

import (
	"errors"
	"fmt"
	"os"
)

// Queue 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array [5]int //数组模拟队列
	front int //表示指向队列首
	rear int //表示指向队列尾部
}

// AddQueue 添加数据到队列
func (que *Queue)AddQueue(val int)(err error){
	if que.IsFull() {
		return errors.New("队列已满")
	}
	que.array[que.rear]=val//将数据放入队列
	que.rear=(que.rear+1)%que.maxSize//移动下标位置
	return
}

// GetQueue 取出一个数据
func (que *Queue)GetQueue()(val int,err error) {
	if que.IsEmpty(){
		return -1,errors.New("队列无数据")
	}
	val=que.array[que.front]
	que.front=(que.front+1)%que.maxSize
	return val,err
}

// ShowQueue 显示队列，找到队首，然后遍历到队尾
func (que *Queue) ShowQueue()  {
	fmt.Println("队列当前的情况是")
	if que.IsEmpty(){
		fmt.Println("没有数据")
	}else {
		//定义一个辅助变量
		testfront:=que.front
		for i:=0;i<que.QuSize();i++{
			fmt.Printf("arr[%d]=%d ",testfront,que.array[testfront])
			testfront=(testfront+1)%que.maxSize
		}
		fmt.Println()
		fmt.Println(que.array)
	}

}

// IsFull 判断队列有没有满
func (que *Queue)IsFull() bool  {
	//留出最后一个位置，如果当前的位置+1取余和头位置想等则满了
	return (que.rear+1)%que.maxSize==que.front
}

// IsEmpty 判断队列为空
func (que *Queue)IsEmpty() bool  {
	//如果头位置和尾部位置相等则没有数据
	return que.front==que.rear
}

// QuSize 判断队列中的数据个数
func (que *Queue)QuSize() int{
	return (que.rear+que.maxSize-que.front)%que.maxSize
}
//编写一个主函数 测试
func main() {
	//先创建一个队列
	queue:=&Queue{
		maxSize: 5,
		front: 0,
		rear: 0,
	}
	var key int
	var val int
	for  {
		fmt.Println("输入1，表示添加数据")
		fmt.Println("输入2，表示从队列中获取数据")
		fmt.Println("输入3，表示显示队列")
		fmt.Println("输入4，表示退出")
		fmt.Scanf("%d",&key)
		switch key {
		case 1:
			fmt.Println("输人要加入队列的数")
			fmt.Scanf("%d",&val)
			err:=queue.AddQueue(val)
			if err==nil {
				fmt.Println("加入队列成功")
			}else{
				fmt.Println(err)
			}
		case 2:
			val,err:=queue.GetQueue()
			if err == nil{
				fmt.Println("从队列中取出一个数",val)
			}else {
				fmt.Println(err)
			}
		case 3:
			queue.ShowQueue()

		case 4:
			os.Exit(0)
		}
	}

}
