package linetable

import (
	"fmt"
)

type List interface{
	Add(elem int)
	Insert(i, elem int)
	Delete(i, elem int)
	Query(elem int) int
	Diskpaly()
	Update(i, elem int) bool
}

type Node struct {
	value int
	Next  *Node
}

type nodeList struct {
	head 	*Node
	length 	int
}

type ListService struct {
	// nodeList nodeList
	list List
}

func NewNodeList() ListService {
	return ListService{
		// nodeList: nodeList{
		// 	head: &Node{},
		// 	length: 0,
		// },
		list: &nodeList{},
	}
} 

func (n *nodeList) Add(elem int) {
	tmp := &Node{}
	tmp = n.head
	for (tmp.Next != nil) {
		tmp = tmp.Next
	}
	tmp.Next = &Node{
		value: elem,
		Next: nil,
	}
	n.length++
}

func (n *nodeList) Insert(i, elem int) {
	tmp := &Node{}
	tmp = n.head
	for (i - 1 > 0 && tmp != nil) {
		tmp = tmp.Next
		i--
	}
	if tmp == nil{
		panic("error")
	}
	if tmp.Next != nil {
		next := tmp.Next
		tmp.Next = &Node{
			value: elem,
			Next: next,
		}
	} else {
		tmp.Next = &Node{
			value: elem,
			Next:nil,
		}
	}
	n.length++
}

func (n *nodeList) Delete(i, elem int) {
	tmp := &Node{}
	tmp = n.head
	for  (i - 1 > 0 && tmp != nil) {
		tmp = tmp.Next
		i--
	}
	if tmp == nil{
		panic("error")
	}
	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}
	n.length--
}

func (n *nodeList) Query(elem int) int {
	tmp := &Node{}
	tmp = n.head
	i := 1
	for  (tmp.Next != nil) {
		if tmp.Next.value == elem {
			return i
		}
		tmp = tmp.Next
		i++
	}
	return -1
}

func (n *nodeList) Diskpaly() {
	tmp := n.head
	for tmp.Next != nil{
		fmt.Printf("%v", tmp.Next.value)
	}
	fmt.Printf("\n")
}

func (n *nodeList) Update(i, elem int) bool {
	tmp := n.head
	k := 1
	for tmp.Next != nil{
		if k == i {
			tmp.Next.value = elem
			return true
		} 
		tmp = tmp.Next
		k++
	}
	return false
}


