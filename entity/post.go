package entity

type Post struct {
	ID      int64  `sql:"id" json:"id"`
	UserId  int64  `sql:"user_id" json:"user_id"`
	Title   string `sql:"title" json:"title"`
	Content string `sql:"content" json:"content"`
}

func (Post) TableName() string {
	return "posts"
}
