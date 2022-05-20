package inmemory

import (
	"errors"
	"simple-library/core/book"
)

var booksMemory = []*book.Book{}

type InMemoryBookRepo struct{}

func (repo *InMemoryBookRepo) GetByCode(code string) (*book.Book, error) {
	var book *book.Book
	for _, bookInMemory := range booksMemory {
		if bookInMemory.Code == code {
			book = bookInMemory
			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (repo *InMemoryBookRepo) GetAll() ([]*book.Book, error) {
	return booksMemory, nil
}

func (repo *InMemoryBookRepo) Save(book *book.Book) error {
	for i, bookInMemory := range booksMemory {
		if bookInMemory.Code == book.Code {
			booksMemory[i] = book

			return nil
		}
	}

	booksMemory = append(booksMemory, book)
	return nil
}

func (repo *InMemoryBookRepo) Delete(code string) error {
	for i, bookInMemory := range booksMemory {
		if bookInMemory.Code == code {
			booksMemory = append(booksMemory[:i], booksMemory[i+1:]...)
			return nil
		}
	}

	return errors.New("book not found")
}
