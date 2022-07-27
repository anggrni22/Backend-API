package registry

import "akrab-bangkit2022-api/controller"

func (r *registry) NewModulAndQuizController() controller.ModulAndQuizController {
	return controller.NewModulAndQuizController(
		r.NewModulAndQuizUsecase(),
		r.NewUserAccess(),
	)
}
func (r *registry) NewUsersController() controller.UserController {
	return controller.NewUserController(r.NewUserUsecase())
}