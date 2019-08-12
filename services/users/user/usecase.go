package user

import (
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
)

type Usecase interface {
	GetUserByID(string) (*_models.Person, error)
}
