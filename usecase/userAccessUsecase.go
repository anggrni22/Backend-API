package usecase

import (
	"crypto/md5"
	"encoding/hex"

	entity "akrab-bangkit2022-api/entity"
	"akrab-bangkit2022-api/repository"
	userReposotory "akrab-bangkit2022-api/repository/user"

	response "akrab-bangkit2022-api/response/user"
	rspUser "akrab-bangkit2022-api/response/user"

	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
)

type userAccess struct {
	BaseRepository repository.BaseRepository
	UserAccessRepository userReposotory.UserAccessRepository
	UserRepository       userReposotory.UsersRepository
}

type UserAccess interface {
	GenerateToken(input entity.UserAccess) string
	ValidToken(token string) bool
	GenerateHashPasswordForLogin(iduser int, password string) string
	GenerateHashPasswordForRegister(salt string, password string) string
}

func NewUserAccess(
	br repository.BaseRepository,
	uar userReposotory.UserAccessRepository,
	ur userReposotory.UsersRepository,
) UserAccess {
	return &userAccess{br, uar, ur}
}

func (uu *userAccess) GenerateToken(input entity.UserAccess) string {
	var token string
	cekAcc, _ := uu.UserAccessRepository.Find(input.IDUser)
	if cekAcc.Token == "" {
		token = uuid.NewV4().String()
		input.Token = token
		err := uu.UserAccessRepository.Create(input)
		if err != nil {
		}
	} else {
		resToken:=response.ResponeRegister{}
		copier.Copy(&resToken, &cekAcc)

		token = resToken.Token
	}
	return token
}
func (uu *userAccess) ValidToken(token string) bool {
	res := uu.UserAccessRepository.ValidToken(token)
	return res
}
func (uu *userAccess) GenerateHashPasswordForLogin(iduser int, password string) string {
	finduser, _ := uu.UserRepository.Find(iduser)
	responseuser := rspUser.UsertResultUser{}
	copier.Copy(&responseuser, &finduser)

	passwordplussalt := password + responseuser.Salt
	passwordmd5 := md5.Sum([]byte(passwordplussalt))
	md5tostring := hex.EncodeToString(passwordmd5[:])
	passwordmd5plussalt := md5tostring + responseuser.Salt
	passwordmd5final := md5.Sum([]byte(passwordmd5plussalt))
	md5tostringfinal := hex.EncodeToString(passwordmd5final[:])

	return md5tostringfinal
}
func (uu *userAccess) GenerateHashPasswordForRegister(salt string, password string) string {
	passwordplussalt := password + salt
	passwordmd5 := md5.Sum([]byte(passwordplussalt))
	md5tostring := hex.EncodeToString(passwordmd5[:])
	passwordmd5plussalt := md5tostring + salt
	passwordmd5final := md5.Sum([]byte(passwordmd5plussalt))
	md5tostringfinal := hex.EncodeToString(passwordmd5final[:])

	return md5tostringfinal
}
