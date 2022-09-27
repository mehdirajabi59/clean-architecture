package entity

type UserUseCase interface {
	FindAll() ([]User, error)
	Find(id int64) (User, error)
	Create(u *User) error
	Delete(u *User) error
}

type PostUseCase interface {
	FindAll() ([]Post, error)
	Find(id int64) (Post, error)
	Create(u *Post) error
	Delete(u *Post) error
}
