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

