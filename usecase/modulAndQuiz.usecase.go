package usecase

import (
	entity "akrab-bangkit2022-api/entity"
	"akrab-bangkit2022-api/repository"
	userRepository "akrab-bangkit2022-api/repository/user"
	res "akrab-bangkit2022-api/response"

	"github.com/jinzhu/copier"
)

type ModulAndQuizUsecase interface {
	ModulAndQuizByLevel(token string, level string)(interface{}, error)
	FindAllLevel(token string)(levels []entity.Level, err error)
	FindAllModul(token string) (modul []entity.Modul, err error)
	FindAllLevelByTipe(token string, tipe string)(levels []entity.Level, err error)
	FindAllModulByTipe(token string, tipe string)(modul []entity.Modul, err error)
}

type modulAndQuizUsecase struct {
	repositoryModul repository.ModulReposotory
	repositoryQuiz repository.QuizReposotory
	repositoryLevel repository.LevelReposotory
	UserAccessRepository userRepository.UserAccessRepository
}

func NewModulAndQuizUsecase(
	m repository.ModulReposotory,
	q repository.QuizReposotory,
	l repository.LevelReposotory,
	uas userRepository.UserAccessRepository,
) ModulAndQuizUsecase {
	return &modulAndQuizUsecase{m, q, l, uas}
}

func (u *modulAndQuizUsecase) ModulAndQuizByLevel(token string, level string)(result interface{}, err error){
	validasitoken := u.UserAccessRepository.ValidToken(token)
	if validasitoken {
		modul, err := u.repositoryModul.FindModulByLevel(level)
		quiz, _ := u.repositoryQuiz.FindQuizByLevel(level)
		levels, _ := u.repositoryLevel.FindLevelByLevel(level)
		
		allres:= res.AllRsp{}
		copier.Copy(&allres, &levels)
		
		allres.Modul=append(allres.Modul, modul...)
		allres.Quiz=append(allres.Quiz, quiz... )
	
	return allres, err

	}
	return nil, err
	
}

func (u *modulAndQuizUsecase) FindAllModul(token string) (modul []entity.Modul, err error){
	validasitoken := u.UserAccessRepository.ValidToken(token)
	if validasitoken {
	return u.repositoryModul.FindAllModul()
	}
	return

}

func (u *modulAndQuizUsecase) FindAllLevel(token string)(levels []entity.Level, err error){
	validasitoken := u.UserAccessRepository.ValidToken(token)
	if validasitoken {
		return u.repositoryLevel.FindAllLevel()
	}
	return
		
}

func (u *modulAndQuizUsecase) FindAllLevelByTipe(token string, tipe string)(levels []entity.Level, err error){
	validasitoken := u.UserAccessRepository.ValidToken(token)
	if validasitoken {
		return u.repositoryLevel.FindAllLevelByTipe(tipe)
	}
	return
		
}

func (u *modulAndQuizUsecase) FindAllModulByTipe(token string, tipe string)(modul []entity.Modul, err error){
	validasitoken := u.UserAccessRepository.ValidToken(token)
	if validasitoken {
		return u.repositoryModul.FindAllModulByTipe(tipe)
	}
	return
		
}