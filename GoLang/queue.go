package main

import (
	"fmt"
	"strconv"
)
type listNode struct {
	value int
	next  *listNode
}

type queue struct {
	first *listNode
	last *listNode
}

func makeQueue() (*queue) {
	return &queue{first:nil,last:nil}
}
func (q queue) isEmpty() bool {
	return( q.first == nil)
}

func (q *queue) enqueue(x int) {
    if q.isEmpty(){
    	q.first = &listNode{value:x, next:q.last}
    }else if q.first!=nil && q.last == nil{
    	q.last = &listNode{value:x, next:nil}
    	q.first.next = q.last
    }else{
    	var n *listNode= &listNode{value:x, next:nil}
    	q.last.next=n
    	q.last=n

    }
}

func (q queue) head() int{
	return q.first.value
}
func (q *queue) dequeue() int{
	var t int = q.head()
    q.first = q.first.next
    return t
}

func (q queue) String() string {
	if q.isEmpty() {
		return "[]"
	} else {
		var s string = "[("
		s += strconv.Itoa(q.head())
		s += ")"
		var n *listNode = q.first.next
		for (n != nil) {
			s += " " + strconv.Itoa(n.value)
			n = n.next
		}
		s += "]"
		return s
	}
}

func main() {
	var S *queue = makeQueue()
	fmt.Println(S)
	fmt.Println("Pushing 1 ...")
	S.enqueue(1)
	fmt.Println(S)
	fmt.Println("Pushing 2 ...")
	S.enqueue(2)
	fmt.Println(S)
	fmt.Println("Pushing 3 ...")
	S.enqueue(3)
	fmt.Println(S)
	fmt.Println("Pushing 4 ...")
	S.enqueue(4)
	fmt.Println(S)
	fmt.Println("Popping off", S.dequeue(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.dequeue(),"...")
	fmt.Println(S)
	fmt.Println("Pushing 5 ...")
	S.enqueue(5)
	fmt.Println(S)
	fmt.Println("Pushing 6 ...")
	S.enqueue(6)
	fmt.Println(S)
	fmt.Println("Popping off", S.dequeue(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.dequeue(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.dequeue(),"...")
	fmt.Println(S)
}	 