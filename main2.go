//go:build ignore

package main

import (
	"fmt"

	pos "github.com/kamildrazkiewicz/go-stanford-nlp"
)

func main() {
	var (
		tagger *pos.Tagger
		res    []*pos.Result
		err    error
	)

	if tagger, err = pos.NewTagger(
		"ext/english-left3words-distsim.tagger",    // path to model
		"ext/stanford-postagger.jar"); err != nil { // path to jar tagger file
		fmt.Print(err)
		return
	}
	if res, err = tagger.Tag("I am Moeka. Nice to meet you."); err != nil {
		fmt.Print(err)
		return
	}
	for _, r := range res {
		fmt.Println(r.Word, r.TAG, r.Description())
	}

}