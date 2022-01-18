package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Top int //栈顶
	arr [6]int
}

//入栈
func (st *Stack) Push(val int) (err error) {
	//判断栈是否已满
	if st.Top == len(st.arr)-1 {
		fmt.Println("栈满")
		return errors.New("栈满")
	}
	//添加数据
	st.Top++
	st.arr[st.Top] = val
	return
}

//出栈
func (st *Stack) Pop() (val int, err error) {
	//判断有没有数据
	if st.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("栈满")
	}
	val = st.arr[st.Top]
	st.Top--
	return val, err
}

//遍历栈
func (st *Stack) Show() {
	if st.Top == -1 {
		fmt.Println("栈空")
		return
	}
	for i := st.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, st.arr[i])
	}
}
func main() {
	st := &Stack{
		Top: -1,
	}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	st.Push(6)
	st.Show()

	fmt.Println()
	val, _ := st.Pop()
	fmt.Println(val)
	val, _ = st.Pop()
	fmt.Println(val)
	val, _ = st.Pop()
	fmt.Println(val)
	val, _ = st.Pop()
	val, _ = st.Pop()
	fmt.Println()

	st.Show()


}
