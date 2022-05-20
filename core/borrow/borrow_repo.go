package borrow

type BorrowRepo interface {
	GetByID(id int) (*Borrow, error)
	GetAll() ([]*Borrow, error)
	Save(borrow *Borrow) error
	Delete(id int) error
}
