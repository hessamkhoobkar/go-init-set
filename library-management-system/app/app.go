package app

import (
	"fmt"
	"library-managment-system/books"
)

func Session() {
    
    var option string

    fmt.Print("================================\nLibrary Management Program Menu:\n1. Add a book\n2. Display books\n3. Sort books by year of publication\n4. Exit\n")
    option = GetUserOption()
    validatedOption := ValidateInput(option)
    for validatedOption == 0 {
		fmt.Println("Please enter a valid number between 1 and 4")
        option = GetUserOption()
        validatedOption = ValidateInput(option)
	}
    // add book
    // sort books
    switch validatedOption {
        case 1:
            fmt.Println("1")
        case 2:
            books.Display()
        case 3:
            fmt.Println("3")  
        case 4:
            Exit()            
    }

}