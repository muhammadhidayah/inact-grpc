package user

import (
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
)

type Usecase interface {
	GetUserByID(string) (*_models.Person, error)
	GetAllUser() (*_models.PersonList, error)
	AddPerson(*_models.Person) error
	UpdatePerson(*_models.Person) error
	DeletePersonByID(string) error
}
