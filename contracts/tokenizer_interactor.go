package contracts

import "github.com/techerfan/ir-project/internal/entity"

//go:generate mockgen -destination=../internal/mocks/interactor/tokenizer.go -package=interactor_mock . Tokenizer

type Tokenizer interface {
	AddDocument(doc entity.Document) error
	FindToken(word string) *entity.Token
	FindDocument(id uint) *entity.Document
	AssignSkipPointers()
}
