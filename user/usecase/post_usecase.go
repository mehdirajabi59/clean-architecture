package usecase

import "clean/entity"

type postUseCase struct {
	repo entity.PostRepository
}

func NewPostUseCase(repo entity.PostRepository) entity.PostUseCase {
	return &postUseCase{
		repo: repo,
	}
}

func (pc *postUseCase) FindAll() ([]entity.Post, error) {
	return pc.repo.FindAll()
}

func (pc *postUseCase) Find(id int64) (entity.Post, error) {
	return pc.repo.FindById(id)
}

func (pc *postUseCase) Create(u *entity.Post) error {
	return pc.repo.Create(u)
}

func (pc *postUseCase) Delete(u *entity.Post) error {
	return pc.repo.Delete(u)
}
