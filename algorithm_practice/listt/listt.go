package listt

import (
	"fmt"
	"reflect"
)

// List 表示一个可以保存任何类型的值的单链表。
type List[T any] struct {
	next *List[T]
	val  T
}

func (o *List[T]) GetLastTwoP() (*List[T], *List[T]) {
	var pre, p *List[T]
	pre, p = nil, o
	for {
		if p.next == nil {
			break
		}
		pre = p
		p = p.next
	}
	return pre, p
}

func (o *List[T]) Add(v T) {
	b := List[T]{nil, v}
	_, p := o.GetLastTwoP()
	p.next = &b
}

func (o *List[T]) DropOne() {
	pre, _ := o.GetLastTwoP()
	if pre == nil {
		fmt.Println("List only has one item, nothing done.")
		return
	}
	pre.next = nil
}

func (o *List[T]) Show() string {
	res := "Start: "
	p := o
	sep := "--->"
	for {
		v := reflect.ValueOf(p.val)
		switch v.Kind() {
		case reflect.Int:
			res += fmt.Sprintf(" %v ", p.val)
		case reflect.Int32:
			res += fmt.Sprintf(" %c ", p.val)
		}

		if p.next == nil {
			res += " :End"
			fmt.Println(res)
			return res
		} else {
			res += sep
			p = p.next
		}
	}
}

func CreateOne[T any](val T) *List[T] {
	return &List[T]{nil, val}
}
