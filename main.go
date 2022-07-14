package main

import (
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"path"

	contracts "github.com/techerfan/ir-project/contracts"
	"github.com/techerfan/ir-project/internal/entity"
	"github.com/techerfan/ir-project/internal/interactor/query"
	"github.com/techerfan/ir-project/internal/interactor/tokenizer"
)

// const sampleText1 = `This is a sample text for testing tokenizer package. to be or not to be.`
// const sampleText2 = `sample text 2 to be or not to be`

func main() {

	pathPtr := flag.String("path", "", "The path of documents.")
	invertedPtr := flag.Bool("inverted", false, "Bool --Enbales inverted index algorithm.")
	skipPtr := flag.Bool("skip", false, "Bool -- Enables skip list algorithm. Type the words after the flag.")
	positionalPtr := flag.Bool("positional", false, "Bool -- Enables positional index algorithm. Type the words after the flag.")
	phrasePtr := flag.String("phrase", "", "The phrase you want to look for by positional indexing.")

	flag.Parse()

	if (*skipPtr && *invertedPtr) ||
		(*invertedPtr && *positionalPtr) ||
		(*positionalPtr && *skipPtr) {
		fmt.Println("You can only select one algotithm at the same time")
		return
	}

	if len(*pathPtr) == 0 {
		fmt.Println("You must enter the path of documents")
		return
	}

	files, err := ioutil.ReadDir(*pathPtr)
	if err != nil {
		fmt.Printf("cannot read the path: %v\n", err)
		return
	}

	t := tokenizer.New()

	for i, f := range files {
		doc := entity.Document{
			Id:   uint(i),
			Path: path.Join(*pathPtr, f.Name()),
		}
		t.AddDocument(doc)
	}

	if *skipPtr {
		words := flag.Args()
		if len(words) == 0 {
			fmt.Println("No word specified")
			return
		}
		q := query.New(t)
		results := q.Intersect(true, words...)
		printResult(results, t)
	} else if *invertedPtr {
		words := flag.Args()
		if len(words) == 0 {
			fmt.Println("No word specified")
			return
		}
		q := query.New(t)
		results := q.Intersect(false, words...)
		printResult(results, t)
	} else if *positionalPtr {
		if len(*phrasePtr) == 0 {
			fmt.Println("No phrase specified")
		}
		q := query.New(t)
		ans := q.SearchPhrase(*phrasePtr)
		if len(ans) < 1 {
			fmt.Println("No result")
			return
		}
		fmt.Println("Results:")
		for _, a := range ans {
			doc := t.FindDocument(a.DocId)
			fmt.Printf("start position: %d, doc: %s\n", a.Position, doc.Path)
		}
	}
}

func printResult(l *list.List, t contracts.Tokenizer) {
	if l != nil && l.Len() > 0 {
		fmt.Println("Results:")
		el := l.Front()
		for el != nil {
			doc := t.FindDocument(el.Value.(*entity.Posting).DocId)
			if doc == nil {
				fmt.Println("Unknown source")
				continue
			}
			fmt.Println(doc.Path)
			el = el.Next()
		}
	} else {
		fmt.Println("No result")
	}
}
