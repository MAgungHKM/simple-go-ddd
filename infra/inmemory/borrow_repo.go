package inmemory

import (
	"errors"
	"simple-library/core/borrow"
)

var borrowsMemory = []*borrow.Borrow{}

type InMemoryBorrowRepo struct{}

func (repo *InMemoryBorrowRepo) GetByID(id int) (*borrow.Borrow, error) {
	var borrow *borrow.Borrow
	for _, borrowInMemory := range borrowsMemory {
		if borrowInMemory.ID == id {
			borrow = borrowInMemory
			return borrow, nil
		}
	}

	return nil, errors.New("borrow not found")
}

func (repo *InMemoryBorrowRepo) GetAll() ([]*borrow.Borrow, error) {
	return borrowsMemory, nil
}

func (repo *InMemoryBorrowRepo) Save(borrow *borrow.Borrow) error {
	for i, borrowInMemory := range borrowsMemory {
		if borrowInMemory.ID == borrow.ID {
			borrowsMemory[i] = borrow

			return nil
		}
	}

	borrowsMemory = append(borrowsMemory, borrow)
	return nil
}

func (repoo *InMemoryBorrowRepo) Delete(id int) error {
	for i, borrowInMemory := range borrowsMemory {
		if borrowInMemory.ID == id {
			borrowsMemory = append(borrowsMemory[:i], borrowsMemory[i+1:]...)
			return nil
		}
	}

	return errors.New("borrow not found")
}
