package algors

import "container/list"

type Bag struct {
	list list.List
}

func NewBag() *Bag {
	return &Bag{
		list: *list.New(),
	}
}

func (b *Bag) IsEmpty() bool {
	return b.list.Len() == 0
}

func (b *Bag) Size() int {
	return b.list.Len()
}

func (b *Bag) Add(item interface{}) {
	b.list.PushBack(item)
}
