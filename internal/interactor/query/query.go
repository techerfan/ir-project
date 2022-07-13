package query

import (
	"container/list"
	"ir-project/contracts"
	"ir-project/internal/entity"
	"ir-project/pkg/trimmer"
	"sort"
	"strings"
)

type (
	Query struct {
		tokenizer contracts.Tokenizer
	}

	PhraseAnswer struct {
		DocId    uint
		Position int
	}
)

func New(t contracts.Tokenizer) *Query {
	return &Query{
		tokenizer: t,
	}
}

func (q *Query) Intersect(withSkips bool, words ...string) *list.List {

	var tokens []*entity.Token

	for _, w := range words {
		token := q.tokenizer.FindToken(trimmer.Trim(w))
		if token == nil {
			return nil
		}
		tokens = append(tokens, token)
	}

	q.sortTokens(tokens)

	result := tokens[0].PostingList
	for i := 1; i < len(tokens); i++ {
		if withSkips {
			result = q.skipPointer(result, tokens[i].PostingList)
		} else {
			result = q.intersection(result, tokens[i].PostingList)
		}
		if result == nil {
			break
		}
	}
	return result
}

func (q *Query) SearchPhrase(phrase string) []PhraseAnswer {
	words := strings.Fields(phrase)

	var tokens []*entity.Token

	for _, w := range words {
		word := trimmer.Trim(w)

		if token := q.tokenizer.FindToken(word); token != nil {
			tokens = append(tokens, token)
		} else {
			return nil
		}
	}

	var finalResults []positionalAnswer
	for i := 1; i < len(tokens); i++ {
		results := q.positionalIndex(i, tokens[0].PostingList, tokens[i].PostingList)

		if len(results) == 0 {
			return nil
		}

		for _, p := range results {
			if p.position1 < p.position2 {
				finalResults = append(finalResults, p)
			}
		}
	}

	var answer []PhraseAnswer

	for _, fr1 := range finalResults {
		var counter = 0
		for j := 1; j < len(tokens); j++ {
			for k, fr2 := range finalResults {
				if fr1.docId == fr2.docId &&
					fr1.position1 == fr2.position1 &&
					fr2.position2 == fr1.position2+j {
					counter++
					finalResults = append(finalResults[:k], finalResults[k+1:]...)
					break
				}
			}
		}
		if counter == len(tokens)-2 {
			answer = append(answer, PhraseAnswer{
				DocId:    fr1.docId,
				Position: fr1.position1,
			})
		}
	}

	return answer
}

// sortTokens sorts tokens by their frequency.
func (t *Query) sortTokens(tokens []*entity.Token) {
	sort.Slice(tokens, func(a, b int) bool {
		return tokens[a].Frequency < tokens[b].Frequency
	})
}
