package main

import "fmt"

type Stack[] string 

func(s *Stack)isEmpty()bool{
	return len(*s)==0
}


func (s *Stack)push(str string){
	*s=append(*s, str)
}


func (s *Stack)pop()(string,bool){
  if  s.isEmpty() {
	  return "",false
  }else{
	last:=len(*s)-1
	element:=(*s)[last]
	*s=(*s)[:last]
	return element,true
  }
   
}


func main(){
	var  myStack Stack
	name,ok:=myStack.pop()
	fmt.Println(name,ok)
	myStack.push("jashfeer")
	myStack.push("jon")
	myStack.push("josaf")
	myStack.push("jinga")
	fmt.Println(myStack)
	name,ok=myStack.pop();
	fmt.Println(myStack)
	fmt.Println(name,ok)


}