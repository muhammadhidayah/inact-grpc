package user

import _models "github.com/muhammadhidayah/inact-grpc/common/models"

type Repository interface {
	GetUserById(string) (*_models.Person, error)
}
