package main

import "fmt"

func main() {
	fmt.Print(" Enter your choice : ")
	choice := ""
	fmt.Scanln(&choice)
	//todo -> which choice -----------------------------------
	if choice == "book" || choice == "Book" {
		book1 := Book{
			title:  "Alvido Vatan",
			author: "A'zam xoshimiy",
			pages:  117,
		}
		fmt.Print("\n---------------------------------")
		fmt.Println(book1.PrintBook())
		fmt.Println("--------------------------------")

		//todo -> CONSTRUCTOR OF THE BOOK CLASS --------------
		a := ConstructorBook("Smth", "Someone", 0)
		fmt.Println("\nAnother book : ", a)
	} else if choice == "phone" || choice == "Phone" {

		//todo -> PHONE CLASS ------------------------------------
		phone1 := CellPhone{
			phoneModel:     "Samsung S 22 ultra",
			phoneColor:     "Saphyr Blue",
			phonePrice:     "$1200",
			phoneProcessor: "Snap Dragon",
			whereReleased:  "USA",
			warantyYear:    1,
		}
		fmt.Println(phone1.PrintPhoneInfo())
	}

}
