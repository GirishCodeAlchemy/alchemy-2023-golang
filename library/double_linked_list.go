package library

import (
	"container/list"
	"fmt"
)

type entry struct {
	key, value int
}

func display_double_linked_list(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d -> ", e.Value)
	}
	fmt.Println()

	// for e := list.Front(); e != nil; e = e.Next() {
	// 	entry := e.Value.(*entry)
	// 	fmt.Printf("%d: %d -> ", entry.key, entry.value)
	// }
}

func insert_before(value int, mark int, curlist *list.List) *list.Element {
	fmt.Println("Insert before")
	display_double_linked_list(curlist)
	var marklist *list.Element
	for e := curlist.Front(); e != nil; e = e.Next() {
		if e.Value == mark {
			marklist = e
			break
		}
	}

	return curlist.InsertBefore(value, marklist)

	// return list.InsertBefore(6, list.Front())
}

func double_linked_list() {
	fmt.Println("Double Linked List")

	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	display_double_linked_list(l)

	l.PushBack(4)
	l.PushFront(5)
	display_double_linked_list(l)

	l.InsertBefore(6, l.Back())
	l.InsertBefore(6, l.Front())
	insert_before(6, 2, l)
	display_double_linked_list(l)

}

func init() {
	double_linked_list()
}
