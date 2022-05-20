package member

type MemberRepo interface {
	GetByCode(code string) (*Member, error)
	GetAll() ([]*Member, error)
	Save(member *Member) error
	Delete(code string) error
}
