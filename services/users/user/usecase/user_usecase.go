package usecase

import (
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
	"github.com/muhammadhidayah/inact-grpc/services/users/user"
)

type userUseCase struct {
	ur user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.Usecase {
	return &userUseCase{
		ur: userRepo,
	}
}

func (uc *userUseCase) GetUserByID(personid string) (*_models.Person, error) {
	res, err := uc.ur.GetUserById(personid)

	if err != nil {
		return nil, err
	}

	return res, nil
}
