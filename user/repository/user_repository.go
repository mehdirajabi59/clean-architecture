package repository

import (
	"clean/entity"
	"database/sql"
	"errors"
)

var user entity.User = entity.User{}
var ErrorUserNotFound error = errors.New("user Not found")

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) entity.UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (ur *mysqlUserRepository) FindAll() ([]entity.User, error) {

	var u []entity.User

	row, err := ur.db.Query("SELECT id, name FROM " + user.TableName())

	if err != nil {
		return u, err
	}

	for row.Next() {
		var oneUser entity.User
		err = row.Scan(&oneUser.ID, &oneUser.Name)

		if err != nil {
			return u, err
		}

		u = append(u, oneUser)
	}

	return u, nil
}

func (ur *mysqlUserRepository) FindById(id int64) (entity.User, error) {

	var u entity.User

	row := ur.db.QueryRow("SELECT id, name FROM "+user.TableName()+" WHERE id=?", id)

	err := row.Scan(&u.ID, &u.Name)

	return u, err
}

func (ur *mysqlUserRepository) Create(u *entity.User) error {

	res, err := ur.db.Exec("INSERT INTO "+user.TableName()+" (name) VALUES(?)", u.Name)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		u.ID = id
	}
	return nil
}

func (ur *mysqlUserRepository) Delete(u *entity.User) error {

	res, err := ur.db.Exec("DELETE FROM "+user.TableName()+" WHERE id=?", u.ID)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrorUserNotFound
	}
	return nil
}
