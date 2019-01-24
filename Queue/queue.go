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
	var q QueueNode
	q.Create(s1,s2,s3)
	q.Print()
	fmt.Println(q.Length())
	s4 := Person{1007, "赵六", "男", 19}
	q.Push(s4)
	q.Print()
	fmt.Println(q.Length())
	n:=q.Length()
	for i:=0; i<n; i++ {
		node:=q.Pop()
		fmt.Println(node)
		q.Print()
		fmt.Println(q.Length())
	}
	q.Push(s3)
	q.Print()
	fmt.Println(q.Length())
	q.Clear()
	fmt.Print(q)
}

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

func (q *QueueNode) Create(data ...interface{}) {
	if q==nil || len(data)==0 {
		return
	}
	node:=q
	for _,v:=range data{
		newNode:=new(QueueNode)
		newNode.Data=v
		node.Next=newNode
		node=node.Next
	}
}

func (q *QueueNode) Print()  {
	if q==nil {
		return
	}
	node:=q.Next
	for node!=nil {
		fmt.Print(node.Data," ")
		node=node.Next
	}
	fmt.Println()
}

func (q *QueueNode) Length() int {
	if q==nil {
		return 0
	}
	node:=q.Next
	n:=0
	for node!=nil {
		n++
		node=node.Next
	}
	return n
}

func (q *QueueNode) Push(data interface{})  {
	if q==nil {
		return
	}
	node:=q
	for node.Next!=nil {
		node=node.Next
	}
	newNode:=new(QueueNode)
	newNode.Data=data
	node.Next=newNode
}

func (q *QueueNode) Pop() interface{} {
	if q==nil {
		return nil
	}
	node:=q.Next
	if node==nil {
		q.Next=nil
	} else {
		q.Next=node.Next
	}
	return node.Data
}

func (q *QueueNode) Clear()  {
	if q==nil {
		return
	}
	q.Next.Clear()
	q.Data=nil
	q.Next=nil
	q=nil
}
