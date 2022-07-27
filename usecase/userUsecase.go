package usecase

import (
	"akrab-bangkit2022-api/config"
	"akrab-bangkit2022-api/entity"
	"akrab-bangkit2022-api/pkg/helper"
	"akrab-bangkit2022-api/repository"
	userRepository "akrab-bangkit2022-api/repository/user"
	req "akrab-bangkit2022-api/request/user"
	rspUser "akrab-bangkit2022-api/response/user"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/thanhpk/randstr"
)

type userUsecase struct {
	BaseRepository       repository.BaseRepository
	UserRepository       userRepository.UsersRepository
	UserAccessRepository userRepository.UserAccessRepository
	UserDataRepository   userRepository.UserDataRepository
	UserAccess           UserAccess
}

type UserUsecase interface{
	Find(id int, token string) (interface{}, error)
	Create(input req.RegRegister) (interface{}, error)
	IsDuplicateEmail(username string) bool
	VerifyEmail(username string, password string) interface{}
}

func NewUserUsecse(
	br repository.BaseRepository,
	ur userRepository.UsersRepository,
	uar userRepository.UserAccessRepository,
	udr userRepository.UserDataRepository,
	ua UserAccess,
) UserUsecase{
	return &userUsecase{
		br,
		ur,
		uar,
		udr,
		ua,
	}
}

func (uu *userUsecase) Find(id int, token string) (interface{}, error) {
	var err error
	findDataUser,_ := uu.UserDataRepository.Find(id)
	findUser, err := uu.UserRepository.Find(id)
	res := rspUser.UsertResultUser{}
	copier.Copy(&res, &findUser)
	copier.Copy(&res, &findDataUser)

	resp := rspUser.ResponeRegister{}
	resp.JwtToken, _ = helper.GenerateJwt(strconv.Itoa(res.IdUser), res.Email,
		config.C.Auth.ApplicationIssuer, config.C.Auth.CmsSecret, config.C.Auth.ExpiredTimeHour)
	resp.User = res
	resp.Token = token
	resp.ServerTime = time.Now().Format("2006-01-02 15:04:05")

	return resp, err
}

func (uu *userUsecase) IsDuplicateEmail(email string) bool {
	filter := map[string]interface{}{
		"email": email,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return true
	}
	if res.IdUser == 0 {
		return true
	}
	return false
}

func (uu *userUsecase) Create(input req.RegRegister) (interface{}, error) {
	uu.BaseRepository.BeginTx()
	salt := randstr.String(32)
	passwordhashgenerator := uu.UserAccess.GenerateHashPasswordForRegister(salt, input.Password)
	superadmin := entity.User{
		Email:        input.Email,
		Password:     passwordhashgenerator,
		Salt:         salt,
	}
	rUser, errrUser := uu.UserRepository.Create(superadmin)
	if errrUser != nil {
		uu.BaseRepository.RollbackTx()
		return false, errrUser
	}
	userdata := entity.UserData{
		IDUser: rUser.IdUser,
		Email:        input.Email,
		FullName:    input.FullName,
		
	}
	rProfileUser, _ := uu.UserDataRepository.Create(userdata)

	access := entity.UserAccess{
		IDUser: rUser.IdUser,
	}

	token := uu.UserAccess.GenerateToken(access)
	resp, err := uu.Find(rProfileUser.IDUser, token)
	if err != nil {
		uu.BaseRepository.RollbackTx()
		return false, err
	}
	return resp, err
}

func (uu *userUsecase) VerifyEmail(email string, password string) interface{} {
	filter := map[string]interface{}{
		"email": email,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return false
	}
	uda := entity.UserAccess{
		IDUser:       res.IdUser,
	}
	token := uu.UserAccess.GenerateToken(uda)
	passwordhasher := uu.UserAccess.GenerateHashPasswordForLogin(res.IdUser, password)
	if res.Email == email && res.Password == passwordhasher {
		resp, _ := uu.Find(res.IdUser, token)
		return resp
	}
	return false
}