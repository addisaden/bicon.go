package base

import "testing"

func TestGetBooks(t *testing.T) {
	b1 := GetBereanInterlinear1()
	b7 := GetBereanInterlinear7()
	b66 := GetBereanInterlinear66()

	these_books := map[string]VerseType{}

	for key, value := range b1 {
		these_books[key] = value
	}

	for key, value := range b7 {
		these_books[key] = value
	}

	for key, value := range b66 {
		these_books[key] = value
	}

	collected_books := GetBooks(1, 7, 66)

	for key, value := range these_books {
		if collected_books[key].book != value.book {
			t.Errorf("In %v the book should be %v", collected_books[key], value.book)
		}
		if collected_books[key].chapter != value.chapter {
			t.Errorf("In %v the chapter should be %v", collected_books[key], value.chapter)
		}
		if collected_books[key].verse != value.verse {
			t.Errorf("In %v the verse should be %v", collected_books[key], value.verse)
		}
		if collected_books[key].text != value.text {
			t.Errorf("In %v the text should be %v", collected_books[key], value.text)
		}
	}

	for key, value := range collected_books {
		if these_books[key].book != value.book {
			t.Errorf("In %v the book should be %v", these_books[key], value.book)
		}
		if these_books[key].chapter != value.chapter {
			t.Errorf("In %v the chapter should be %v", these_books[key], value.chapter)
		}
		if these_books[key].verse != value.verse {
			t.Errorf("In %v the verse should be %v", these_books[key], value.verse)
		}
		if these_books[key].text != value.text {
			t.Errorf("In %v the text should be %v", these_books[key], value.text)
		}
	}
}
