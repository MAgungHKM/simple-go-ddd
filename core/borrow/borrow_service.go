package borrow

import (
	"errors"
	"fmt"
	"simple-library/core/book"
)

type BorrowService struct {
	borrowRepo BorrowRepo
	bookRepo   book.BookRepo
}

func (borrowService *BorrowService) New(borrowRepo BorrowRepo, bookRepo book.BookRepo) {
	borrowService.borrowRepo = borrowRepo
	borrowService.bookRepo = bookRepo
}

func (borrowService *BorrowService) Apply(borrow *Borrow) error {
	err := isBorrowedBooksValid(borrowService.bookRepo, borrow.Books)
	if err != nil {
		return err
	}

	// TODO: Should be transactional with try catch
	// if any of the following steps fails, rollback
	err = borrowService.borrowRepo.Save(borrow)
	if err != nil {
		return err
	}

	for _, book := range borrow.Books {
		book.Quantity -= 1
		err = borrowService.bookRepo.Save(book)
		if err != nil {
			return err
		}
	}

	return nil
}

func isBorrowedBooksValid(bookRepo book.BookRepo, books []*book.Book) error {
	for bookIndex, book := range books {
		// begin check quantity
		bookInMemory, err := bookRepo.GetByCode(book.Code)
		if err != nil {
			return err
		}

		if bookInMemory.Quantity-1 < 0 {
			return fmt.Errorf("book #%s quantity is less than 0", book.Code)
		}
		// end check quantity

		// begin check duplicate
		for nextBookIndex, nextBook := range books {
			if bookIndex == nextBookIndex {
				continue
			}

			if book.Code == nextBook.Code {
				return errors.New("book code is duplicated")
			}
		}
		// end check duplicate
	}

	return nil
}
