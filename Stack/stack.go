package main

import "fmt"

type Person struct {
	Id int
	Name string
	Sex string
	Age int
}

func main()  {
	s1 := Person{1001, "张三", "男", 18}
	s2 := Person{1008, "李四", "男", 20}
	s3 := Person{1010, "王五", "男", 24}
	var s StackNode
	s.Create(s1,s2,s3)
	s.Print()
	fmt.Println(s.Length())
	s4 := Person{1007, "赵六", "男", 19}
	s.Push(s4)
	s.Print()
	fmt.Println(s.Length())
	n:=s.Length()
	for i:=0; i<n; i++ {
		node:=s.Pop()
		fmt.Println(node)
		s.Print()
		fmt.Println(s.Length())
	}
	s.Push(s1)
	s.Print()
	fmt.Println(s.Length())
	s.Clear()
	fmt.Print(s)
}

type StackNode struct {
	Data interface{}
	Next *StackNode
}

func (s *StackNode) Create(data ...interface{}) {
	if s==nil || len(data)==0 {
		return
	}
	for _,v:=range data{
		newNode:=new(StackNode)
		newNode.Data=v
		newNode.Next=s.Next
		s.Next=newNode
	}
}

func (s *StackNode) Print()  {
	if s==nil {
		return
	}
	node:=s.Next
	for node!=nil {
		fmt.Print(node.Data," ")
		node=node.Next
	}
	fmt.Println()
}

func (s *StackNode) Length() int {
	if s==nil {
		return 0
	}
	node:=s.Next
	n:=0
	for node!=nil {
		n++
		node=node.Next
	}
	return n
}

func (s *StackNode) Push(data interface{}) {
	if s==nil {
		return
	}
	newNode:=new(StackNode)
	newNode.Data=data
	newNode.Next=s.Next
	s.Next=newNode
}

func (s *StackNode) Pop() interface{} {
	if s==nil {
		return nil
	}
	node:=s.Next
	s.Next=node.Next
	return node.Data
}

func (s *StackNode) Clear()  {
	if s==nil {
		return
	}
	s.Next.Clear()
	s.Data=nil
	s.Next=nil
	s=nil
}
