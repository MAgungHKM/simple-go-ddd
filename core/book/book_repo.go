package book

type BookRepo interface {
	GetByCode(code string) (*Book, error)
	GetAll() ([]*Book, error)
	Save(book *Book) error
	Delete(code string) error
}
