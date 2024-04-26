package prac

import (
	"fmt"
	"sync"
)

type listItem struct {
	val  interface{}
	prev *listItem
	next *listItem
}

type list struct {
	depth uint
	first *listItem
	last  *listItem
	sync.RWMutex
}

func NewList() *list {
	return &list{}
}

func (l *list) insert(val interface{}) {
	l.Lock()
	defer l.Unlock()
	data := &listItem{val: val, prev: l.last, next: l.first}
	if l.depth == 0 {
		l.first = data
		l.last = data
		l.depth++
		return
	}
	l.first.prev = data
	l.last.next = data
	l.first = data
	l.depth++

}

func (l *list) find(val interface{}) bool {
	if l.depth == 0 {
		return false
	}
	data := l.first
	n := l.depth
	for n > 0 {

		if data.val == val {
			return true
		}
		data = data.next
		n--
	}
	return false
}

func (l *list) delete(val interface{}) error {
	l.RLock()
	if l.depth == 0 {
		return fmt.Errorf("can not delete val %v, for the list depth is 0", val)
	}

	data := l.first
	n := l.depth
	l.RUnlock()
	l.Lock()
	defer l.Unlock()
	for n > 0 {

		if data.val == val {
			data.next.prev = data.prev
			data.prev.next = data.next
			data.val = nil
			data.next = nil
			data.prev = nil
			l.depth--
			return nil

		}
		data = data.next
		n--
	}
	return fmt.Errorf("can not delete val %v, for can not find the item in the list", val)

}

func (l *list) print() {
	data := l.first
	n := l.depth
	for n > 0 {

		fmt.Println(data.val)
		data = data.next
		n--
	}
}
