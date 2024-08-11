package base

import (
	"fmt"
	"slices"
)

func all_books() []string {
	all_books := []string{
		"",
		"genesis", "exodus", "leviticus", "numbers", "deuteronomy", "joshua", "judges", "ruth", "i_samuel", "ii_samuel", "i_kings", "ii_kings", "i_chronicles", "ii_chronicles", "ezra", "nehemiah", "esther", "job", "psalms", "proverbs", "ecclesiastes", "song_of_solomon", "isaiah", "jeremiah", "lamentations", "ezekiel", "daniel", "hosea", "joel", "amos", "obadiah", "jonah", "micah", "nahum", "habakkuk", "zephaniah", "haggai", "zechariah", "malachi", "matthew", "mark", "luke", "john", "acts", "romans", "i_corinthians", "ii_corinthians", "galatians", "ephesians", "philippians", "colossians", "i_thessalonians", "ii_thessalonians", "i_timothy", "ii_timothy", "titus", "philemon", "hebrews", "james", "i_peter", "ii_peter", "i_john", "ii_john", "iii_john", "jude", "revelation_of_john",
	}
	return all_books
}

func GetBookName(bookIndex int) (string, error) {
	books := all_books()
	if bookIndex > len(books) || bookIndex < 1 {
		return "", fmt.Errorf("could not found a book with index %v", bookIndex)
	}
	return books[bookIndex], nil
}

func GetBookIndex(bookName string) (int, error) {
	books := all_books()
	bookIndex := slices.Index(books, bookName)
	if bookIndex >= 1 {
		return bookIndex, nil
	}
	return 0, fmt.Errorf("could not found an index for the book %v", bookName)
}
