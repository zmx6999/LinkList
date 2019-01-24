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
	fmt.Println(node.Length())
	s6 := Person{1000, "陈八", "女", 30}
	node.Insert(1,s6)
	node.Print()
	fmt.Println(node.Length())
	fmt.Println(node.GetData(3))
	node.Delete(1)
	node.Print()
	fmt.Println(node.Length())
	fmt.Println(node.Search(s1))
	node.DeleteByData(s1)
	node.Print()
	fmt.Println(node.Length())
	node.Destroy()
	fmt.Print(node)
}

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func (node *LinkNode) Create(data ...interface{})  {
	if node==nil || len(data)==0 {
		return
	}
	head:=node
	_node:=node
	for _,v:=range data{
		newNode:=new(LinkNode)
		newNode.Data=v
		_node.Next=newNode
		_node=_node.Next
	}
	_node.Next=head.Next
}

func (node *LinkNode) Print()  {
	if node==nil || node.Next==nil {
		return
	}
	start:=node.Next
	_node:=node.Next
	for _node!=nil {
		fmt.Print(_node.Data," ")
		_node=_node.Next
		if _node==start {
			fmt.Print(_node.Data)
			break
		}
	}
	fmt.Println()
}

func (node *LinkNode) Length() int {
	if node==nil || node.Next==nil {
		return 0
	}
	n:=0
	start:=node.Next
	_node:=node.Next
	for _node!=nil {
		n++
		_node=_node.Next
		if _node==start {
			break
		}
	}
	return n
}

func (node *LinkNode) Insert(i int,data interface{})  {
	if node==nil {
		return
	}
	if i<=0 || i>node.Length()+1 {
		return
	}
	j:=0
	preNode:=node
	_node:=node
	for j<i && _node!=nil {
		j++
		preNode=_node
		_node=_node.Next
	}
	newNode:=new(LinkNode)
	newNode.Data=data
	newNode.Next=_node
	preNode.Next=newNode
	if i==1 {
		start:=newNode.Next
		tmpNode:=newNode
		for tmpNode.Next!=nil {
			tmpNode=tmpNode.Next
			if tmpNode.Next==start {
				break
			}
		}
		tmpNode.Next=newNode
	}
}

func (node *LinkNode) GetData(i int) interface{} {
	if node==nil || node.Next==nil {
		return nil
	}
	if i<=0 || i>node.Length() {
		return nil
	}
	j:=0
	_node:=node
	for j<i && _node!=nil {
		j++
		_node=_node.Next
	}
	return _node.Data
}

func (node *LinkNode) Delete(i int)  {
	if node==nil || node.Next==nil {
		return
	}
	if i<=0 || i>node.Length() {
		return
	}
	if node.Length()==1 {
		_node:=node.Next
		node.Next=nil
		_node.Data=nil
		_node.Next=nil
		_node=nil
		return
	}
	j:=0
	preNode:=node
	_node:=node
	for j<i && _node!=nil {
		j++
		preNode=_node
		_node=_node.Next
	}
	preNode.Next=_node.Next
	if i==1 {
		start:=_node
		tmpNode:=_node
		for tmpNode.Next!=nil {
			tmpNode=tmpNode.Next
			if tmpNode.Next==start {
				break
			}
		}
		tmpNode.Next=start.Next
	}
	_node.Data=nil
	_node.Next=nil
	_node=nil
}

func (node *LinkNode) Search(data interface{}) int {
	if node==nil || node.Next==nil {
		return -1
	}
	i:=0
	_node:=node
	for _node!=nil {
		i++
		if i>node.Length() {
			break
		}
		_node=_node.Next
		if _node!=nil && reflect.TypeOf(data)==reflect.TypeOf(_node.Data) && data==_node.Data {
			return i
		}
	}
	return -1
}

func (node *LinkNode) DeleteByData(data interface{})  {
	if node==nil || node.Next==nil {
		return
	}
	i:=0
	preNode:=node
	_node:=node
	for _node!=nil {
		i++
		if i>node.Length() {
			return
		}
		preNode=_node
		_node=_node.Next
		if _node!=nil && reflect.TypeOf(data)==reflect.TypeOf(_node.Data) && data==_node.Data {
			if node.Length()==1 {
				_node:=node.Next
				node.Next=nil
				_node.Data=nil
				_node.Next=nil
				_node=nil
				return
			}
			break
		}
	}
	preNode.Next=_node.Next
	if i==1 {
		start:=_node
		tmpNode:=_node
		for tmpNode.Next!=nil {
			tmpNode=tmpNode.Next
			if tmpNode.Next==start {
				break
			}
		}
		tmpNode.Next=start.Next
	}
	_node.Data=nil
	_node.Next=nil
	_node=nil
}

func (node *LinkNode) Destroy()  {
	if node==nil {
		return
	}
	start:=node.Next
	tmpNode:=node
	for tmpNode.Next!=nil {
		tmpNode=tmpNode.Next
		if tmpNode.Next==start {
			tmpNode.Next=nil
		}
	}
	node.destroy()
}

func (node *LinkNode) destroy()  {
	if node==nil {
		return
	}
	node.Next.destroy()
	node.Data=nil
	node.Next=nil
	node=nil
}
