package entity

type User struct {
	ID   int64  `sql:"id" json:"id"`
	Name string `sql:"name" json:"name"`
}

func (u User) TableName() string {

	return "users"
}
