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

func (uc *userUseCase) AddPerson(data *_models.Person) error {
	err := uc.ur.AddPerson(data)
	return err
}

func (uc *userUseCase) UpdatePerson(data *_models.Person) error {
	err := uc.ur.UpdatePerson(data)
	return err
}

func (uc *userUseCase) GetAllUser() (*_models.PersonList, error) {
	res, err := uc.ur.GetAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *userUseCase) DeletePersonByID(id string) error {
	err := uc.ur.DeletePersonByID(id)

	return err
}
