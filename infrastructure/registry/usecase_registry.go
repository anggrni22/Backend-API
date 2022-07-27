package registry

import (
	"akrab-bangkit2022-api/repository"
	userRepository "akrab-bangkit2022-api/repository/user"
	"akrab-bangkit2022-api/usecase"
	userusecase "akrab-bangkit2022-api/usecase"
)

func (r *registry) NewModulAndQuizUsecase() usecase.ModulAndQuizUsecase{
	return usecase.NewModulAndQuizUsecase(
		repository.NewModulRepository(r.db),
		repository.NewQuizRepository(r.db),
		repository.NewLevelRepository(r.db),
		userRepository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
	)
}
func (r *registry) NewUserUsecase() userusecase.UserUsecase {
	return userusecase.NewUserUsecse(
		repository.NewBaseRepository(r.db),
		userRepository.NewUsersRepository(repository.NewBaseRepository(r.db)),
		userRepository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		userRepository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		userusecase.NewUserAccess(
			repository.NewBaseRepository(r.db),
			userRepository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
			userRepository.NewUsersRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
func (r *registry) NewUserAccess() userusecase.UserAccess {
	return userusecase.NewUserAccess(
		repository.NewBaseRepository(r.db),
		userRepository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		userRepository.NewUsersRepository(repository.NewBaseRepository(r.db)),
	)
}