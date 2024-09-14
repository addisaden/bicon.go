package search

import "github.com/addisaden/bicon.go/base"

type ResultType struct {
	Query    string
	Limit    int
	Offset   int
	Context  base.VerseTextType
	Selected []string
	Results  []ResultVerseType
}

type ResultVerseType struct {
	Verse base.VerseTextType
	Match []int
}
