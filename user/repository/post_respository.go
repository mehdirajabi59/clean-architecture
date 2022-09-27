package repository

import (
	"clean/entity"
	"database/sql"
)

var post entity.Post = entity.Post{}

type mysqlPostRepository struct {
	db *sql.DB
}

func NewMysqlPostRepository(db *sql.DB) entity.PostRepository {
	return &mysqlPostRepository{
		db: db,
	}
}

func (pr *mysqlPostRepository) FindAll() ([]entity.Post, error) {

	var p []entity.Post

	row, err := pr.db.Query("SELECT id, user_id, title, content FROM " + post.TableName())

	if err != nil {
		return p, err
	}

	for row.Next() {
		var onePost entity.Post
		err = row.Scan(&onePost.ID, &onePost.UserId, &onePost.Title, &onePost.Content)

		if err != nil {
			return p, err
		}

		p = append(p, onePost)
	}

	return p, nil
}

func (pr *mysqlPostRepository) FindById(id int64) (entity.Post, error) {

	var p entity.Post

	row := pr.db.QueryRow("SELECTid, user_id, title, content FROM "+post.TableName()+" WHERE id=?", id)

	err := row.Scan(&p.ID, &p.UserId, &p.Title, &p.Content)

	return p, err
}

func (pr *mysqlPostRepository) Create(p *entity.Post) error {

	res, err := pr.db.Exec("INSERT INTO "+post.TableName()+" (title, user_id, content) VALUES(?)", p.Title, p.UserId, p.Content)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		p.ID = id
	}
	return nil
}

func (pr *mysqlPostRepository) Delete(u *entity.Post) error {

	_, err := pr.db.Exec("DELETE FROM "+post.TableName()+" WHERE id=?", u.ID)

	return err
}
