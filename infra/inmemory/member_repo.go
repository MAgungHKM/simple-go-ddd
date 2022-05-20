package inmemory

import (
	"errors"
	"simple-library/core/member"
)

var membersMemory = []*member.Member{}

type InMemoryMemberRepo struct{}

func (repo *InMemoryMemberRepo) GetByCode(code string) (*member.Member, error) {
	var member *member.Member
	for _, memberInMemory := range membersMemory {
		if memberInMemory.Code == code {
			member = memberInMemory
			return member, nil
		}
	}

	return nil, errors.New("member not found")
}

func (repo *InMemoryMemberRepo) GetAll() ([]*member.Member, error) {
	return membersMemory, nil
}

func (repo *InMemoryMemberRepo) Save(member *member.Member) error {
	for i, memberInMemory := range membersMemory {
		if memberInMemory.Code == member.Code {
			membersMemory[i] = member

			return nil
		}
	}

	membersMemory = append(membersMemory, member)
	return nil
}

func (repo *InMemoryMemberRepo) Delete(code string) error {
	for i, memberInMemory := range membersMemory {
		if memberInMemory.Code == code {
			membersMemory = append(membersMemory[:i], membersMemory[i+1:]...)
			return nil
		}
	}

	return errors.New("member not found")
}
