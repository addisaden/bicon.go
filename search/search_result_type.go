package search

import "github.com/addisaden/bicon.go/base"

type ResultType struct {
	Query   string
	Limit   int
	Offset  int
	Results []base.VerseTextType
}
