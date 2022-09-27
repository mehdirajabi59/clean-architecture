package entity


type UserRepository interface{
	FindAll() ([]User, error)
	FindById(id int64) (User, error)
	Create(u *User) error
	Delete(u *User) error
}


type PostRepository interface{
	FindAll() ([]Post, error)
	FindById(id int64) (Post, error)
	Create(u *Post) error
	Delete(u *Post) error
}