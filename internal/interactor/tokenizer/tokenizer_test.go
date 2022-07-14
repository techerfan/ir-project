package tokenizer

import (
	"container/list"
	"os"
	"testing"

	"github.com/techerfan/ir-project/internal/entity"
)

const sampleText1 = `This is a sample text for testing tokenizer package. to be or not to be.`
const sampleText2 = `sample text 2`

func setupTestEnvironment(t *testing.T) (*os.File, *os.File, func() error) {
	dir, err := os.MkdirTemp("", "trends_test_*")
	if err != nil {
		t.Error(err)
	}

	file1, err := os.CreateTemp(dir, "sample_doc_*")
	if err != nil {
		t.Error(err)
	}
	if _, err = file1.Write([]byte(sampleText1)); err != nil {
		t.Error(err)
	}

	file2, err := os.CreateTemp(dir, "sample_doc_*")
	if err != nil {
		t.Error(err)
	}
	if _, err = file2.Write([]byte(sampleText2)); err != nil {
		t.Error(err)
	}

	return file1, file2, func() error {
		if err := file1.Close(); err != nil {
			return err
		}
		if err = file2.Close(); err != nil {
			return err
		}
		err = os.RemoveAll(dir)
		return err
	}
}

func TestAddDocument(t *testing.T) {
	file1, file2, teardown := setupTestEnvironment(t)
	defer teardown()

	doc1 := entity.Document{
		Id:   0,
		Path: file1.Name(),
	}

	doc2 := entity.Document{
		Id:   1,
		Path: file2.Name(),
	}

	tokenizer := New()
	tokenizer.AddDocument(doc1)
	tokenizer.AddDocument(doc2)

	if len(tokenizer.tokens) < 1 {
		t.Fail()
	}

	for i := range tokenizer.tokens {
		for j := i + 1; j < len(tokenizer.tokens); j++ {
			if tokenizer.tokens[i] == tokenizer.tokens[j] {
				t.Errorf("repeated token: %s", tokenizer.tokens[i].Word)
			}
		}
	}
}

func TestAssignSkipPointers(t *testing.T) {
	file1, file2, teardown := setupTestEnvironment(t)
	defer teardown()

	doc1 := entity.Document{
		Id:   0,
		Path: file1.Name(),
	}

	doc2 := entity.Document{
		Id:   1,
		Path: file2.Name(),
	}

	tokenizer := New()
	tokenizer.AddDocument(doc1)
	tokenizer.AddDocument(doc2)

	postingList := list.New().Init()
	postingList.PushFront(&entity.Posting{DocId: 0})
	postingList.PushFront(&entity.Posting{DocId: 1})
	postingList.PushFront(&entity.Posting{DocId: 2})
	postingList.PushFront(&entity.Posting{DocId: 3})
	postingList.PushFront(&entity.Posting{DocId: 4})
	postingList.PushFront(&entity.Posting{DocId: 5})
	postingList.PushFront(&entity.Posting{DocId: 6})
	postingList.PushFront(&entity.Posting{DocId: 7})
	postingList.PushFront(&entity.Posting{DocId: 8})
	postingList.PushFront(&entity.Posting{DocId: 9})

	tokenizer.tokens = []*entity.Token{
		{
			Word:        "sample",
			PostingList: postingList,
		},
	}

	tokenizer.AssignSkipPointers()
	counter := 0
	for p := tokenizer.tokens[0].PostingList.Front(); p != nil; p = p.Next() {
		if p.Value.(*entity.Posting).SkipPointer != nil {
			counter++
		}
	}

	if counter != 7 {
		t.Errorf("expected: %d, found: %d", 7, counter)
	}
}
