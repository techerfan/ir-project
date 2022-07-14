package query

import (
	"container/list"

	"github.com/techerfan/ir-project/internal/entity"
)

func (q *Query) skipPointer(p1, p2 *list.List) *list.List {
	answer := list.New().Init()
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
			if ps1.SkipPointer != nil && ps1.SkipPointer.Value.(*entity.Posting).DocId <= ps2.DocId {
				for skip := el1.Value.(*entity.Posting).SkipPointer; skip != nil && (skip.Value.(*entity.Posting).DocId <= ps2.DocId); el1 = skip {
				}
			} else {
				el1 = el1.Prev()
			}
		} else if ps2.SkipPointer != nil && ps2.SkipPointer.Value.(*entity.Posting).DocId <= ps1.DocId {
			for skip := el2.Value.(*entity.Posting).SkipPointer; skip != nil && (skip.Value.(*entity.Posting).DocId <= ps1.DocId); el2 = skip {
			}
		} else {
			el2 = el2.Prev()
		}
	}
	return answer
}
