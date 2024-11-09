/* queue_test.go - Test for generic queue implementation
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

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue *queue = NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for i := 1; i < 6; i++ {
		item := queue.Dequeue()

		if item != i {
			t.Error("TestQueue failed...", i, item)
		}
	}
}

func TestStack(t *testing.T) {
	var stack *stack = newStack()

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	for i := 4; i > 0; i-- {
		item := stack.pop()

		if item != i {
			t.Error("TestStack failed...", i, item)
		}
	}
}

func TestList(t *testing.T) {
	var list *list = NewList()

	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.print()
	fmt.Println(list.find(1))
	fmt.Println(list.find(5))
	list.delete(1)
	list.print()
}

func TestArrStack(t *testing.T) {
	s := "23*54*+9-"
	res := PostfixCompute0428(s)
	fmt.Println(res)
	
}


func TestInfixToPostfix(t *testing.T) {
	//	abc*+ea*+f+
	//s := "a+b*c+e*a+f"
	//	ab+c*d-e*
	s := "((a+b)*c-d)*e"
	res := infixToPostfixP(s)
	fmt.Println(res)
	
}

func TestReverseLinkedList(t *testing.T) {
	l := newLinkedList(2)
	l = l.insert(3)
	l = l.insert(4)
	l = l.insert(5)
	l = l.insert(6)
	l.print()
	l = l.reverse()
	l.print()
}



func TestItemQueue(t *testing.T) {
	q := newItemQueue()
	q.enqueue(1)
	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(6)
	q.enqueue(4)
	q.dequeue()
	fmt.Println(q.getFront())
	fmt.Println(q.getRear())
}


func TestArrQueue(t *testing.T) {
	q := newArrQueue(7)
	q.enqueue(1)
	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(6)
	q.enqueue(4)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	//q.dequeue()
	q.enqueue(1)
	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(6)
	q.enqueue(4)
	fmt.Println(q.getFront())
	fmt.Println(q.getRear())
}


