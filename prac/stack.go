package prac

import (
	"fmt"
	"strconv"
)

type stackItem struct {
	item interface{}
	next *stackItem
}

type stack struct {
	depth uint
	top   *stackItem
}

func newStack() *stack {
	return new(stack)
}

func (s *stack) push(item interface{}) {
	data := &stackItem{item: item, next: s.top}
	s.top = data
	s.depth++
}
func (s *stack) pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	data := s.top
	s.top = data.next
	s.depth--
	return data.item

}
func (s *stack) peek() interface{} {
	return s.top
}

// type ArrStack []string

type ArrStack struct {
	top  uint
	data []string
}

func NewArrStack() *ArrStack {
	res := new(ArrStack)
	return res
}

func (arr *ArrStack) push(s string) {
	arr.data = append(arr.data, s)
	arr.top += 1
}

func (arr *ArrStack) pop() string {
	if arr.top == 0 {
		panic("can not pop empty stack")
	}

	res := arr.data[arr.top-1]
	arr.top -= 1
	// what if [0:0]
	if arr.top == 0 {
		arr.data = []string{}
	} else {
		arr.data = arr.data[0:arr.top]
	}
	return res
}

func (arr *ArrStack) peek() string {
	return arr.data[arr.top-1]
}


func (arr *ArrStack) isEmpty() bool {
	return arr.top == 0
}


var operantMap = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
	"(": true,
	")": true,
}
func PostfixCompute0428(s string) int {
	//operantList := []string{"+","-","*"}
	
	arrStack := NewArrStack()
	for _, item := range s {
		tmp := string(item)
		if _, ok := operantMap[tmp]; !ok {
			arrStack.push(tmp)
		} else {
			//LIFO
			operator2, _ := strconv.Atoi(arrStack.pop())
			operator1, _ := strconv.Atoi(arrStack.pop())
			res := 0
			if tmp == "*" {
				res = operator1 * operator2
			} else if tmp == "+" {
				res = operator1 + operator2
			} else if tmp == "-" {
				res = operator1 - operator2
			}
			resStr := strconv.Itoa(res)
			arrStack.push(resStr)
		}
	}
	result, _ := strconv.Atoi(arrStack.pop())
	return result
}

type linkedList struct {
	val  int
	next *linkedList
}

func newLinkedList(x int) *linkedList {
	res := new(linkedList)
	res.val = x
	return res
}

func (l *linkedList) insert(x int) *linkedList {
	newLinkedList := newLinkedList(x)
	newLinkedList.next = l
	return newLinkedList
}
func (l *linkedList) print() {
	tmp := l
	for tmp != nil {
		fmt.Print(tmp.val)
		fmt.Print(" ")
		tmp = tmp.next
	}
	fmt.Println("--------------")
}

func (l *linkedList) reverse() *linkedList{
	stack := newStack()
	tmp := l
	for tmp != nil {
		stack.push(tmp)
		tmp = tmp.next
	}
	l = stack.pop().(*linkedList)
	tempReverse := l
	continueSign := true
	for continueSign {
		tmp := stack.pop()
		if tmp != nil {
			node := tmp.(*linkedList)
			tempReverse.next = node
			tempReverse = node
		} else {
			tempReverse.next = nil
			continueSign = false
		}
	}
	return l
}

func infixToPostfix(s string) string {

	arrStack := NewArrStack()
	res := ""
	for _, item := range s {
		tmp := string(item)
		if _, ok := operantMap[tmp]; !ok{
			res += tmp
		}else {
			if !arrStack.isEmpty()&&metHigherOperant(tmp,arrStack.peek()){
				//res += tmp
				for !arrStack.isEmpty(){
					res += arrStack.pop()
				}
			}
			arrStack.push(tmp)
		}
	}
	for !arrStack.isEmpty(){
		res += arrStack.pop()
	}
	return res
}


func infixToPostfixP(s string) string {

	arrStack := NewArrStack()
	res := ""
	for _, item := range s {
		tmp := string(item)
		if _, ok := operantMap[tmp]; !ok{
			res += tmp
		}else {
			if !arrStack.isEmpty()&&metHigherOperant(tmp,arrStack.peek()){
				// if met higher; then pop & add till empty or Open
				for !arrStack.isEmpty()&&!metOpenP(arrStack.peek()){
					res += arrStack.pop()
				}
			}

			if metCloseP(tmp) {
				// if ), then pop&add till (, then pop (
				for !metOpenP(arrStack.peek()) {
					res += arrStack.pop()
				}
				// pop out the openP
				arrStack.pop()
			}else {
				// if ( or operant, then push
				arrStack.push(tmp)
			}
		}
	}
	for !arrStack.isEmpty(){
		res += arrStack.pop()
	}
	return res
}

func metHigherOperant(a,b string) bool{
	if (a == "+" || a == "-") && (b == "*"||b=="/") {
		return true
	}
	return false
}

func metOpenP(a string) bool {
	return a == "("
}


func metCloseP(a string) bool {
	return a == ")"
}