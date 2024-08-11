package base

import "testing"

func TestGetBookName(t *testing.T) {
	bookIndexA := 1
	bookNameA, err := GetBookName(bookIndexA)
	wantBookNameA := "genesis"
	if wantBookNameA != bookNameA || err != nil {
		t.Fatalf(`GetBookName(%v) = %v, %v, want match for %v, nil`, bookIndexA, bookNameA, err, wantBookNameA)
	}

	bookIndexB := 300
	bookNameB, err := GetBookName(bookIndexB)
	if err == nil {
		t.Fatalf(`GetBookName(%v) = %v, %v, want an error`, bookIndexB, bookNameB, err)
	}

	bookIndexC := 66
	bookNameC, err := GetBookName(bookIndexC)
	wantBookNameC := "revelation_of_john"
	if wantBookNameC != bookNameC || err != nil {
		t.Fatalf(`GetBookName(%v) = %v, %v, want match for %v, nil`, bookIndexC, bookNameC, err, wantBookNameC)
	}
}
