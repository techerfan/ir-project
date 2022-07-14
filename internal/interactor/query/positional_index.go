package query

import (
	"container/list"
	"math"

	"github.com/techerfan/ir-project/internal/entity"
)

type positionalAnswer struct {
	docId                uint
	position1, position2 int
}

func (q *Query) positionalIndex(k int, p1, p2 *list.List) []positionalAnswer {
	var answer []positionalAnswer
	el1 := p1.Back()
	el2 := p2.Back()

	for el1 != nil && el2 != nil {
		ps1 := el1.Value.(*entity.Posting)
		ps2 := el2.Value.(*entity.Posting)

		if ps1.DocId == ps2.DocId {
			var l []int
			pp1 := ps1.Positions.Back()
			pp2 := ps2.Positions.Back()
			for pp1 != nil {
				for pp2 != nil {
					if (math.Abs(float64(pp1.Value.(int) - pp2.Value.(int)))) <= float64(k) {
						l = append(l, pp2.Value.(int))
					} else if pp1.Value.(int) < pp2.Value.(int) {
						break
					}
					pp2 = pp2.Prev()
				}

				for len(l) != 0 && math.Abs(float64(l[0]-pp1.Value.(int))) > float64(k) {
					l = l[1:]
				}
				for _, lp := range l {
					answer = append(answer, positionalAnswer{
						docId:     ps1.DocId,
						position1: pp1.Value.(int),
						position2: lp,
					})
				}
				pp1 = pp1.Prev()
			}
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
