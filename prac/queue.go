/* queue.go - A generic queue implementation
 *
 * This file is part of ds library.
 *
 * The MIT License (MIT)
 * Copyright (c) <2016> Alexander Kuleshov <kuleshovmail@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
 * of the Software, and to permit persons to whom the Software is furnished to do
 * so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 * INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
 * PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
 * CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
 
package prac

type queue struct {
	current *queueItem
	last *queueItem
	depth int
}

type queueItem struct {
	item interface{}
	prev *queueItem
}

func NewQueue() *queue{
	return &queue{}
}

func (q *queue) Enqueue(item interface{}) {
	data := &queueItem{item:item,prev:nil}
	if q.depth == 0 {
		q.current = data
		q.last = q.current
		q.depth++
		return 
	}
	q.last.prev = data
	q.last = data
	q.depth++
	
}

func (q *queue) Dequeue() interface{}{
	if q.depth == 0 {
		return nil
	}
	data := q.current
	q.current = data.prev
	q.depth--
	return data.item
}

type item struct{
	data int
	next *item
}

type itemQueue struct{
	depth uint
	front *item
	rear *item
}
// O(1)
func newItemQueue() *itemQueue {
	return new(itemQueue)
}

func (q *itemQueue) enqueue(i int) {
	// insert to the front
	q.depth++
	node := item{
		data:i,
	}
	if q.rear == nil{
		q.front = &node
		q.rear = &node
		return
	}
	q.rear.next = &node
	q.rear = &node
}
func (q *itemQueue) dequeue() {
	if q.front == nil{
		panic("queue is empty")
	}
	q.front = q.front.next
	q.depth--
}

func (q *itemQueue) getFront() *item{
	return q.front
}

func (q *itemQueue) getRear() *item{
	return q.rear
}


func (q *itemQueue) isEmpty() bool{
	return q.depth == 0
}

type arrQueue struct {
	front int
	rear int
	depth uint
	max uint
	data []int
	
}

func newArrQueue(max uint) *arrQueue {
	res := new(arrQueue)
	res.max = max
	res.front = -1
	res.rear = -1
	res.data = make([]int, max)
	return res
}



func (q *arrQueue) enqueue(i int) {
	// insert to the front
	if q.depth + 1 > q.max{
		panic("overflow the queue")
	}
	if q.front == -1 {
		q.front++
		q.rear++
		q.depth++
		return
	}
	q.rear = (q.rear+1)%int(q.max)
	// q.front = (q.rear-int(q.depth))
	// if q.front <0 {
	// 	q.front += int(q.max)
	// }
	q.data[q.rear] = i
	q.depth++
	
}
func (q *arrQueue) dequeue() {
	if q.front == -1 {
		panic("can not dequeue from empty queue")
	}
	q.depth--
	if q.depth == 0{
		q.front = -1
		q.rear = -1
		return
	}
	q.front = (q.rear-int(q.depth)+1)
	if q.front <0 {
		q.front += int(q.max)
	}
	

}

func (q *arrQueue) getFront() int{
	return q.data[q.front]
}

func (q *arrQueue) getRear() int{
	return q.data[q.rear]
}


func (q *arrQueue) isEmpty() bool{
	return q.depth == 0
}