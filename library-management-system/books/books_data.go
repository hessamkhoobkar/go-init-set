package books

type Book struct {
	title, author, publication string
}

var Books = []Book{
	{
		title:       "To Kill a Mockingbird",
		author:      "Harper Lee",
		publication: "1960",
	},
}