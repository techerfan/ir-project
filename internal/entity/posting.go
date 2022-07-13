package entity

import "container/list"

type Posting struct {
	DocId       uint
	Positions   *list.List
	SkipPointer *list.Element
}
