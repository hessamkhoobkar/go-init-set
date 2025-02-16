package books

import "fmt"

func Display() {
	for _, book := range Books {
		fmt.Printf("Title: %s, Author: %s, Year of Publication: %s", book.title, book.author, book.publication)
	}
}