package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_models "github.com/muhammadhidayah/inact-grpc/common/models"
	"github.com/muhammadhidayah/inact-grpc/services/users/user"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) user.Repository {
	return &userRepository{Conn}
}

func (ur *userRepository) GetUserById(id string) (*_models.Person, error) {
	rows, err := ur.Conn.Query("SELECT * FROM ts_org_person WHERE person_id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() {
		errClose := rows.Close()

		if errClose != nil {
			fmt.Println(errClose.Error())
		}
	}()

	result := make([]*_models.Person, 0)

	for rows.Next() {
		t := new(_models.Person)
		rows.Scan(
			&t.PersonNo,
			&t.PersonId,
			&t.PersonName,
			&t.PersonPosition,
			&t.PersonMail,
			&t.PersonTelephone,
			&t.PersonPassword,
		)

		result = append(result, t)
	}

	if len(result) > 0 {
		return result[0], nil
	} else {
		return nil, errors.New("Your requested item not found")
	}
}
