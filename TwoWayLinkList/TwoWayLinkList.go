package main

import (
	"fmt"
	"reflect"
)

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
	s4 := Person{1007, "赵六", "男", 19}
	s5 := Person{1666, "刘七", "男", 21}
	var node LinkNode
	node.Create(s1,s2,s3,s4,s5)
	node.Print()
	node.ReversePrint()
	fmt.Println(node.Length())
	s6 := Person{1000, "陈八", "女", 30}
	node.Append(s6)
	node.Print()
	node.ReversePrint()
	fmt.Println(node.Length())
	s7:=Person{1888,"徐莉","女",28}
	node.Insert(6,s7)
	node.Print()
	node.ReversePrint()
	fmt.Println(node.Length())
	fmt.Println(node.GetData(7))
	node.Delete(3)
	node.Print()
	node.ReversePrint()
	fmt.Println(node.Length())
	fmt.Println(node.Search(s6))
	node.DeleteByData(s7)
	node.Print()
	node.ReversePrint()
	fmt.Println(node.Length())
	node.Destroy()
	fmt.Print(node)
}

type LinkNode struct {
	Data interface{}
	Next *LinkNode
	Prev *LinkNode
}

func (node *LinkNode) Create(data ...interface{})  {
	if node==nil {
		return
	}
	_node:=node
	for _,v:=range data{
		newNode:=new(LinkNode)
		newNode.Data=v
		newNode.Next=nil
		newNode.Prev=_node
		_node.Next=newNode
		_node=_node.Next
	}
}

func (node *LinkNode) Print()  {
	if node==nil {
		return
	}
	_node:=node.Next
	for _node!=nil {
		fmt.Print(_node.Data," ")
		_node=_node.Next
	}
	fmt.Println()
}

func (node *LinkNode) ReversePrint()  {
	if node==nil {
		return
	}
	head:=node
	_node:=node
	for _node.Next!=nil {
		_node=_node.Next
	}
	for _node!=nil && _node!=head {
		fmt.Print(_node.Data," ")
		_node=_node.Prev
	}
	fmt.Println()
}

func (node *LinkNode) Length() int {
	if node==nil {
		return 0
	}
	n:=0
	_node:=node.Next
	for _node!=nil {
		n++
		_node=_node.Next
	}
	return n
}

func (node *LinkNode) Append(data interface{})  {
	if node==nil {
		return
	}
	_node:=node
	for _node.Next!=nil {
		_node=_node.Next
	}
	newNode:=new(LinkNode)
	newNode.Data=data
	newNode.Next=nil
	newNode.Prev=_node
	_node.Next=newNode
}

func (node *LinkNode) Insert(i int,data interface{})  {
	if node==nil {
		return
	}
	if i<=0 {
		return
	}
	j:=0
	prenode:=node
	_node:=node
	for j<i && _node!=nil {
		j++
		prenode=_node
		_node=_node.Next
	}
	if _node==nil {
		return
	}
	newNode:=new(LinkNode)
	newNode.Data=data
	newNode.Next=_node
	newNode.Prev=prenode
	prenode.Next=newNode
	_node.Prev=newNode
}

func (node *LinkNode) GetData(i int) interface{} {
	if node==nil {
		return nil
	}
	if i<=0 {
		return nil
	}
	j:=0
	_node:=node
	for j<i && _node!=nil {
		j++
		_node=_node.Next
	}
	if _node==nil {
		return nil
	}
	return _node.Data
}

func (node *LinkNode) Delete(i int)  {
	if node==nil {
		return
	}
	if i<=0 {
		return
	}
	j:=0
	prenode:=node
	_node:=node
	for j<i && _node!=nil {
		j++
		prenode=_node
		_node=_node.Next
	}
	if _node==nil {
		return
	}
	prenode.Next=_node.Next
	if _node.Next!=nil {
		_node.Next.Prev=prenode
	}
	_node.Data=nil
	_node.Next=nil
	_node.Prev=nil
	_node=nil
}

func (node *LinkNode) Search(data interface{}) int {
	if node==nil {
		return -1
	}
	i:=0
	_node:=node
	for _node!=nil {
		i++
		_node=_node.Next
		if _node!=nil && reflect.TypeOf(data)==reflect.TypeOf(_node.Data) && data==_node.Data {
			return i
		}
	}
	return -1
}

func (node *LinkNode) DeleteByData(data interface{})  {
	if node==nil {
		return
	}
	prenode:=node
	_node:=node
	for _node!=nil {
		prenode=_node
		_node=_node.Next
		if _node!=nil && reflect.TypeOf(data)==reflect.TypeOf(_node.Data) && data==_node.Data {
			break
		}
	}
	if _node==nil {
		return
	}
	prenode.Next=_node.Next
	if _node.Next!=nil {
		_node.Next.Prev=prenode
	}
	_node.Data=nil
	_node.Next=nil
	_node.Prev=nil
	_node=nil
}

func (node *LinkNode) Destroy()  {
	if node==nil {
		return
	}
	node.Next.Destroy()
	node.Data=nil
	node.Next=nil
	node.Prev=nil
	node=nil
}
