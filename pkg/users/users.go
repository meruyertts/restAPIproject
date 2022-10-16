package users

import (
	"ts/db"
)

type Users struct {
	Id   int
	Data string
}

func NewUser(id int, data string) *Users {
	return &Users{
		Id:   id,
		Data: data,
	}
}

func (u *Users) Create() error {
	_, err := db.DB.NamedExec(`INSERT INTO users(id, data) VALUES(:id, :data)`,
		map[string]interface{}{
			"id":   u.Id,
			"data": u.Data,
		})

	return err
}

func (u *Users) Update(newData string) error {
	_, err := db.DB.NamedExec(`UPDATE users SET data=:data WHERE id=:id`,
		map[string]interface{}{
			"id":   u.Id,
			"data": newData,
		})
	return err
}

func (u *Users) Read() (*Users, error) {
	rows, err := db.DB.NamedQuery(`SELECT * FROM users WHERE id=:id`, map[string]interface{}{"id": u.Id})
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (u *Users) Delete() error {
	_, err := db.DB.NamedExec(`DELETE FROM users WHERE id=:id`,
		map[string]interface{}{
			"id": u.Id,
		})
	return err
}
