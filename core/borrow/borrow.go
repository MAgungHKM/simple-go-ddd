package borrow

import (
	"simple-library/core/book"
	"simple-library/core/member"
	"simple-library/core/status"

	"time"
)

var autoIncrement int = 1

type Borrow struct {
	ID         int
	Date       time.Time
	ReturnDate time.Time
	// LatestReturnDate time.Time
	Status string

	Books  []*book.Book
	Member *member.Member
}

func (borrow *Borrow) AddBook(book *book.Book) {
	borrow.Books = append(borrow.Books, book)
}

func (borrow *Borrow) DeleteLatestBook() {
	borrow.Books = borrow.Books[:len(borrow.Books)-1]
}

func (borrow *Borrow) New() {
	borrow.ID = autoIncrement
	borrow.Date = time.Now()
	borrow.ReturnDate = time.Now().AddDate(0, 0, 7)
	// borrow.LatestReturnDate = time.Now().AddDate(0, 0, 7)
	borrow.Status = status.Active

	autoIncrement += 1
}
