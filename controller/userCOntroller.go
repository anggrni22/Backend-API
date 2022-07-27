package controller

import (
	"akrab-bangkit2022-api/entities/response"
	"akrab-bangkit2022-api/usecase"
	"net/http"

	req "akrab-bangkit2022-api/request/user"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase usecase.UserUsecase
}
type UserController interface {
	Register(c *gin.Context)
	LoginEmail(c *gin.Context)
}
func NewUserController(
	u usecase.UserUsecase,
	) UserController {
		return &userController{
			u,
		}
	}
func (uc *userController) Register(c *gin.Context) {
	var input req.RegRegister
	var result interface{}
	
	var errRes error

	if err := c.ShouldBind(&input); err != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrShouldBindError,
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if errVal := input.Validate(); errVal != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrFailedToSave,
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
		if !uc.UserUsecase.IsDuplicateEmail(input.Email) {
			resp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrEmailAlreadyUsed,
				},
			}
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		result, errRes = uc.UserUsecase.Create(input)
	

	if errRes != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrEmailAlreadyUsed,
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
		},
		Data: result,
	}
	c.JSON(http.StatusOK, resp)
}

func (uc *userController) LoginEmail(c *gin.Context) {
	var input req.RegLoginEmail

	if err := c.ShouldBind(&input); err != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrShouldBindError,
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := input.Validate(); err != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrFailedToSave,
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result := uc.UserUsecase.VerifyEmail(input.Email, input.Password)
	if result == false {
				resp := response.Response{
					Meta: response.Meta{
						Message:   response.RespMeta.TelErrFailedToLogin,
					},
				}
				c.JSON(http.StatusBadRequest, resp)
				return
			} else {
				resp := response.Response{
					Meta: response.Meta{
						Message:   response.RespMeta.Success,
					},
					Data: result,
				}
				c.JSON(http.StatusOK, resp)
			}
}

