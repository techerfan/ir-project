package query

import (
	"container/list"
	"ir-project/internal/entity"
)

func (q *Query) intersection(p1, p2 *list.List) *list.List {
	answer := list.New().Init()
	// el stands for element
	el1 := p1.Back()
	el2 := p2.Back()
	for el1 != nil && el2 != nil {
		ps1 := el1.Value.(*entity.Posting)
		ps2 := el2.Value.(*entity.Posting)

		if ps1.DocId == ps2.DocId {
			answer.PushFront(ps1)
			el1 = el1.Prev()
			el2 = el2.Prev()
		} else if ps1.DocId < ps2.DocId {
			el1 = el1.Prev()
		} else {
			el2 = el2.Prev()
		}
	}
	return answer
}
