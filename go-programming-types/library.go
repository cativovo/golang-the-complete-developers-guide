//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		returnTime := ""

		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended", book.lended)
	}

	fmt.Println()
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book", title, "not found in library")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("Book", title, "not found from", member.name)
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book", title, "not found in library")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{
		checkOut: time.Now(),
	}

	return true
}

func registerMember(library *Library, name Name) {
	member := Member{
		name:  name,
		books: make(map[Title]LendAudit),
	}

	library.members[name] = member
}

func main() {
	library := Library{
		members: make(map[Name]Member),
		books: map[Title]BookEntry{
			"Book 1": {
				total: 5,
			},
			"Book 2": {
				total: 3,
			},
			"Book 3": {
				total: 6,
			},
			"Book 4": {
				total: 9,
			},
		},
	}

	registerMember(&library, "John Doe")
	registerMember(&library, "Alice Wonderland")
	registerMember(&library, "Foo Bar")

	fmt.Println("Intial")
	printMemberAudits(&library)
	printLibraryBooks(&library)

	member := library.members["Foo Bar"]
	checkedOut := checkoutBook(&library, "Book 1", &member)

	fmt.Println("Check out a book:")
	if checkedOut {
		printMemberAudits(&library)
		printLibraryBooks(&library)
	}

	returned := returnBook(&library, "Book 1", &member)

	fmt.Println("Check in a book:")
	if returned {
		printMemberAudits(&library)
		printLibraryBooks(&library)
	}
}
