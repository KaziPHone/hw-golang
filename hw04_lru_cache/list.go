package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	Key   Key
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushBack(v interface{}) *ListItem {

	item := &ListItem{
		Value: v,
		Prev:  l.last,
	}

	if l.last == nil {
		l.first = item
	} else {
		l.last.Next = item
	}

	l.last = item
	l.len++
	return item
}

func (l *list) PushFront(v interface{}) *ListItem {

	item := &ListItem{
		Value: v,
		Next:  l.first,
	}

	if l.first == nil {
		l.last = item
	} else {
		item.Next.Prev = item
	}

	l.first = item
	l.len++
	return item
}

func (l *list) Len() int {
	return l.len
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.first {
		return
	}

	first := l.first
	next := i.Next
	prev := i.Prev

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	i.Next = first
	i.Prev = nil
	l.first = i
	first.Prev = i
}

func (l *list) Remove(i *ListItem) {

	if i.Prev == nil {
		l.first = i.Next
		l.first.Prev = nil
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.last = i.Prev
		l.last.Next = nil
	} else {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func NewList() List {
	return new(list)
}
