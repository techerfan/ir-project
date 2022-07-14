package tokenizer

import (
	"container/list"
	"io/ioutil"
	"math"
	"strings"

	"github.com/techerfan/ir-project/internal/entity"
	"github.com/techerfan/ir-project/pkg/trimmer"
)

type Tokenizer struct {
	documents []entity.Document
	tokens    []*entity.Token
}

func New() *Tokenizer {
	list.New()
	return &Tokenizer{
		tokens: make([]*entity.Token, 0),
	}
}

// َAddDocument simply reads the specified document from its path
// and find its tokens.
func (t *Tokenizer) AddDocument(doc entity.Document) error {
	t.documents = append(t.documents, doc)

	file, err := ioutil.ReadFile(doc.Path)
	if err != nil {
		return err
	}

	words := strings.Fields(string(file))

	for position, word := range words {
		word = trimmer.Trim(word)
		token := t.FindToken(word)
		// var positions *list.List
		if token == nil {
			token = &entity.Token{
				Word:        word,
				Frequency:   1,
				PostingList: list.New().Init(),
			}
			positions := list.New().Init()
			positions.PushFront(position)
			token.PostingList.PushFront(&entity.Posting{
				DocId:     doc.Id,
				Positions: positions,
			})
			t.tokens = append(t.tokens, token)
		} else if token != nil {
			token.Frequency++
			var positions *list.List
			var posting *list.Element
			for posting = token.PostingList.Front(); posting != nil; posting = posting.Next() {
				val := posting.Value.(*entity.Posting)
				if val.DocId == doc.Id {
					positions = posting.Value.(*entity.Posting).Positions
					break
				}
			}
			// if posting was nil, it means posting related to this
			// doc-id does not exist so that we must create it.
			if posting == nil {
				positions = list.New().Init()
				token.PostingList.PushFront(&entity.Posting{
					DocId:     doc.Id,
					Positions: positions,
				})
			}
			positions.PushFront(position)
		}
	}
	return nil
}

// FindToken tries to find wanted word in the list of tokens
// and returns nil if it cannot find it.
func (t *Tokenizer) FindToken(word string) *entity.Token {
	for _, token := range t.tokens {
		if token.Word == word {
			return token
		}
	}
	return nil
}

func (t *Tokenizer) FindDocument(id uint) *entity.Document {
	for _, doc := range t.documents {
		if doc.Id == id {
			return &doc
		}
	}
	return nil
}

// GetTokens return extracted tokens from documents
func (t *Tokenizer) GetTokens() []*entity.Token {
	return t.tokens
}

// AssignSkipPointers calculates skip pointer for each
// posting and convert the list to a skip list
func (t *Tokenizer) AssignSkipPointers() {
	for _, token := range t.tokens {
		sqrt := math.Sqrt(float64(token.PostingList.Len()))
		length := math.Floor(sqrt)
	Postings:
		for p := token.PostingList.Front(); p != nil; p = p.Next() {
			var next *list.Element = p
			for i := 0; i < int(length); i++ {
				next = next.Next()
				if next == nil {
					break Postings
				}
			}
			p.Value.(*entity.Posting).SkipPointer = next
		}
	}
}
