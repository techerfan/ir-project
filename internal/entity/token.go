package entity

import "container/list"

type Token struct {
	Word        string
	Frequency   uint64
	PostingList *list.List
}
