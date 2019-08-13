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

func (ur *userRepository) AddPerson(data *_models.Person) error {
	stmt, err := ur.Conn.Prepare("INSERT INTO ts_org_person(person_id, person_name, person_position, person_mail, person_telephone, person_password) VALUES($1,$2,$3,$4,$5,$6)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		data.PersonId,
		data.PersonName,
		data.PersonPosition,
		data.PersonMail,
		data.PersonTelephone,
		data.PersonPassword)

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UpdatePerson(data *_models.Person) error {
	stmt, err := ur.Conn.Prepare("UPDATE ts_org_person SET person_id = $1, person_name = $2, person_position = $3, person_mail = $4, person_telephone = $5, person_password = $6 WHERE person_no = $7")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(
		data.PersonId,
		data.PersonName,
		data.PersonPosition,
		data.PersonMail,
		data.PersonTelephone,
		data.PersonPassword,
		data.PersonNo,
	)

	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected > 0 {
		return nil
	} else {
		return errors.New("No Rows Affected")
	}
}

func (ur *userRepository) GetAllUser() (*_models.PersonList, error) {
	rows, err := ur.Conn.Query("SELECT * FROM ts_org_person")

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()

		if err != nil {
			fmt.Println(err.Error())
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

	data := _models.PersonList{
		List: result,
	}

	return &data, nil
}

func (ur *userRepository) DeletePersonByID(id string) error {
	stmt, err := ur.Conn.Prepare("DELETE from ts_org_person WHERE person_id = $1")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected > 0 {
		return nil
	} else {
		return errors.New("No Rows Affected")
	}
}
