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

	bookIndexD := 0
	bookNameD, err := GetBookName(bookIndexD)
	if err == nil {
		t.Fatalf(`GetDookName(%v) = %v, %v, want an error`, bookIndexD, bookNameD, err)
	}
}

func TestGetBookIndex(t *testing.T) {
	bookNameA := "genesis"
	bookIndexA, err := GetBookIndex(bookNameA)
	if bookIndexA != 1 || err != nil {
		t.Errorf("GetDookIndex(%q) = %v, %v", bookNameA, bookIndexA, err)
	}

	bookNameB := "revelation_of_john"
	bookIndexB, err := GetBookIndex(bookNameB)
	if bookIndexB != 66 || err != nil {
		t.Errorf("GetDookIndex(%q) = %v, %v", bookNameB, bookIndexB, err)
	}

	bookNameC := "exodus"
	bookIndexC, err := GetBookIndex(bookNameC)
	if bookIndexC != 2 || err != nil {
		t.Errorf("GetDookIndex(%q) = %v, %v", bookNameC, bookIndexC, err)
	}

	bookNameD := "non_book"
	_, err = GetBookIndex(bookNameD)
	if err == nil {
		t.Errorf("GetDookIndex(%q) should have an error", bookNameD)
	}

	// Test the special case "" - The empty string is in the slice on position 0.
	bookNameE := ""
	_, err = GetBookIndex(bookNameE)
	if err == nil {
		t.Errorf("GetDookIndex(%q) should have an error", bookNameD)
	}
}
