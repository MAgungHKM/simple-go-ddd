package main

import (
	"fmt"
	"simple-library/core/book"
	"simple-library/core/borrow"
	"simple-library/core/member"
	"simple-library/infra/inmemory"
)

func main() {
	var err error

	bookRepo := inmemory.InMemoryBookRepo{}
	borrowRepo := inmemory.InMemoryBorrowRepo{}
	memberRepo := inmemory.InMemoryMemberRepo{}

	harpot1 := book.Book{
		Code:     "HARPOT1",
		ISBN:     "978-4-7741-8411-1",
		Title:    "Harry Potter Volume I",
		Quantity: 2,
	}

	harpot2 := book.Book{
		Code:     "HARPOT2",
		ISBN:     "978-4-7741-8411-2",
		Title:    "Harry Potter Volume II",
		Quantity: 1,
	}

	err = bookRepo.Save(&harpot1)
	if err != nil {
		println(err.Error())
	}

	err = bookRepo.Save(&harpot2)
	if err != nil {
		println(err.Error())
	}

	budi := member.Member{
		Code: "001",
		Name: "Budi",
	}

	ani := member.Member{
		Code: "002",
		Name: "Ani",
	}

	err = memberRepo.Save(&budi)
	if err != nil {
		println(err.Error())
	}

	err = memberRepo.Save(&ani)
	if err != nil {
		println(err.Error())
	}

	budiBorrow := borrow.Borrow{}
	budiBorrow.New()
	budiBorrow.Member = &budi
	budiBorrow.AddBook(&harpot1)
	budiBorrow.AddBook(&harpot2)

	aniBorrow := borrow.Borrow{}
	aniBorrow.New()
	aniBorrow.Member = &ani
	aniBorrow.AddBook(&harpot1)
	aniBorrow.AddBook(&harpot1)

	printMembers(&memberRepo)
	printBooks(&bookRepo)

	borrowService := borrow.BorrowService{}
	borrowService.New(&borrowRepo, &bookRepo)

	// this borrow should be saved
	err = borrowService.Apply(&budiBorrow)
	if err != nil {
		println(err.Error())
	}

	printBorrows(&borrowRepo)
	printBooks(&bookRepo)

	// while this borrow should be failed
	err = borrowService.Apply(&aniBorrow) // error: book is duplicated
	if err != nil {
		println(err.Error())
	}
	println("there should be no new records in borrows & no decrease in quantity of books")

	printBorrows(&borrowRepo)
	printBooks(&bookRepo)

	// let's try again by deleting the duplicate
	aniBorrow.DeleteLatestBook()
	err = borrowService.Apply(&aniBorrow)
	if err != nil {
		println(err.Error())
	}

	printBorrows(&borrowRepo)
	printBooks(&bookRepo)

	// trying to borrow book that is not available
	err = borrowService.Apply(&aniBorrow)
	if err != nil {
		println(err.Error())
	}
}

func printBooks(bookRepo book.BookRepo) {
	booksMemory, _ := bookRepo.GetAll()
	fmt.Println("======================= BOOKS =======================")
	for _, book := range booksMemory {
		fmt.Printf("\nCode: %s\nISBN: %s\nTitle: %s\nQuantity: %d\n\n", book.Code, book.ISBN, book.Title, book.Quantity)
	}
	fmt.Println("======================= BOOKS =======================")

	fmt.Println()
}

func printMembers(memberRepo member.MemberRepo) {
	membersMemory, _ := memberRepo.GetAll()
	fmt.Println("====================== MEMBERS ======================")
	for _, member := range membersMemory {
		fmt.Printf("\nCode: %s\nName: %s\n\n", member.Code, member.Name)
	}
	fmt.Println("====================== MEMBERS ======================")

	fmt.Println()
}

func printBorrows(borrowRepo borrow.BorrowRepo) {
	borrowsMemory, _ := borrowRepo.GetAll()
	fmt.Println("====================== BORROWS ======================")
	for _, borrow := range borrowsMemory {
		fmt.Printf("\nID: %d\nDate: %s\nBorrower: %s\nList of Books:\n", borrow.ID, borrow.Date.Format("2006-01-02"), borrow.Member.Name)
		for _, book := range borrow.Books {
			fmt.Printf("- Title: %s\n", book.Title)
		}
		fmt.Println()
	}
	fmt.Println("====================== BORROWS ======================")

	fmt.Println()
}
