package main

import "fmt"

type stack struct{
	items[]int
}
func (s *stack) Push(i int){
	s.items = append(s.items,i)
}

func(s *stack) Pop()int{
	l:=len(s.items)-1
	 toRemove :=s.items[l]
	 s.items=s.items[:l]
	 return toRemove
}
func main(){
	myStack:=stack{}
	fmt.Println(myStack)
	myStack.Push(20)
	myStack.Push(60)
	myStack.Push(10)
	myStack.Push(30)
	myStack.Push(40)
	myStack.Push(50)
	fmt.Println(myStack)
	value:=myStack.Pop()
	fmt.Println(myStack)
	fmt.Println(value)
	

}