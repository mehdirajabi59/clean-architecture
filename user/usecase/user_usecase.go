package usecase

import "clean/entity"

type userUseCase struct {
	userRepository entity.UserRepository
}

func NewUserUseCase(repo entity.UserRepository) entity.UserUseCase {
	return &userUseCase{
		userRepository: repo,
	}
}

func (uc *userUseCase) FindAll() ([]entity.User, error) {
	return uc.userRepository.FindAll()
}

func (uc *userUseCase) Find(id int64) (entity.User, error) {
	return uc.userRepository.FindById(id)
}

func (uc *userUseCase) Create(u *entity.User) error {
	return uc.userRepository.Create(u)
}

func (uc *userUseCase) Delete(u *entity.User) error {
	return uc.userRepository.Delete(u)
}
