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
		if collected_books[key].Book != value.Book {
			t.Errorf("In %v the book should be %v", collected_books[key], value.Book)
		}
		if collected_books[key].Chapter != value.Chapter {
			t.Errorf("In %v the chapter should be %v", collected_books[key], value.Chapter)
		}
		if collected_books[key].Verse != value.Verse {
			t.Errorf("In %v the verse should be %v", collected_books[key], value.Verse)
		}
		if collected_books[key].Text != value.Text {
			t.Errorf("In %v the text should be %v", collected_books[key], value.Text)
		}
	}

	for key, value := range collected_books {
		if these_books[key].Book != value.Book {
			t.Errorf("In %v the book should be %v", these_books[key], value.Book)
		}
		if these_books[key].Chapter != value.Chapter {
			t.Errorf("In %v the chapter should be %v", these_books[key], value.Chapter)
		}
		if these_books[key].Verse != value.Verse {
			t.Errorf("In %v the verse should be %v", these_books[key], value.Verse)
		}
		if these_books[key].Text != value.Text {
			t.Errorf("In %v the text should be %v", these_books[key], value.Text)
		}
	}
}
