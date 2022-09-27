package registry

import (
	"clean/entity"
	"clean/user/delivery/http"
	"clean/user/repository"
	"clean/user/usecase"
)

func (r *Registry) NewUserRegistry() http.UserHandler {
	return r.newUserHandler()
}

func (r *Registry) newUserHandler() http.UserHandler {
	return http.NewUserHandler(r.newUserUsecase())
}

func (r *Registry) newUserUsecase() entity.UserUseCase {
	return usecase.NewUserUseCase(r.newUserRepository())
}

func (r *Registry) newUserRepository() entity.UserRepository {
	return repository.NewMysqlUserRepository(r.db)
}
