package controller

import (
	"akrab-bangkit2022-api/entities/response"
	"akrab-bangkit2022-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModulAndQuizController interface {
	FindAllModulAndQuizByLevel(c *gin.Context)
	FindAllLevel(c *gin.Context)
	FindAllModul(c *gin.Context)
	FindAllLevelByTipe(c *gin.Context)
	FindAllModulByTipe(c *gin.Context)
}

type modulAndQuizController struct{
	ModulAndQuizUsecase usecase.ModulAndQuizUsecase
	userAcces usecase.UserAccess
}

func NewModulAndQuizController(
	p usecase.ModulAndQuizUsecase,
	u usecase.UserAccess,
	) ModulAndQuizController {
	return &modulAndQuizController{
		p,
	 	u,
	}
}

func (h *modulAndQuizController) FindAllLevel(c *gin.Context){
	paramtoken := c.Query("token")
	result, err := h.ModulAndQuizUsecase.FindAllLevel(paramtoken)
	requestid, _ := c.Get("RequestID")
	validtoken := h.userAcces.ValidToken(paramtoken)

	if validtoken == false {
		validtoken := h.userAcces.ValidToken(paramtoken)
		if validtoken == false {
			rsp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
	}
	
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrItemNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

func (h *modulAndQuizController) FindAllModul(c *gin.Context){
	paramtoken := c.Query("token")
	result, err := h.ModulAndQuizUsecase.FindAllModul(paramtoken)
	requestid, _ := c.Get("RequestID")
	validtoken := h.userAcces.ValidToken(paramtoken)

	if validtoken == false {
		validtoken := h.userAcces.ValidToken(paramtoken)
		if validtoken == false {
			rsp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
	}
	
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrItemNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}


func (h *modulAndQuizController) FindAllModulAndQuizByLevel(c *gin.Context) {
	paramlevel:= c.Param("level")
	paramtoken := c.Query("token")
	requestid, _ := c.Get("RequestID")
	validtoken := h.userAcces.ValidToken(paramtoken)

	if validtoken == false {
		validtoken := h.userAcces.ValidToken(paramtoken)
		if validtoken == false {
			rsp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
	}

	result, err := h.ModulAndQuizUsecase.ModulAndQuizByLevel(paramtoken, paramlevel)

	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrItemNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

func (h *modulAndQuizController) FindAllLevelByTipe(c *gin.Context){
	paramtipe := c.Param("tipe")
	paramtoken := c.Query("token")
	requestid, _ := c.Get("RequestID")
	validtoken := h.userAcces.ValidToken(paramtoken)

	if validtoken == false {
		validtoken := h.userAcces.ValidToken(paramtoken)
		if validtoken == false {
			rsp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
	}

	result, err := h.ModulAndQuizUsecase.FindAllLevelByTipe(paramtoken, paramtipe)
	
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrItemNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

func (h *modulAndQuizController) FindAllModulByTipe(c *gin.Context){
	paramtipe := c.Param("tipe")
	paramtoken := c.Query("token")
	requestid, _ := c.Get("RequestID")
	validtoken := h.userAcces.ValidToken(paramtoken)

	if validtoken == false {
		validtoken := h.userAcces.ValidToken(paramtoken)
		if validtoken == false {
			rsp := response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
	}

	result, err := h.ModulAndQuizUsecase.FindAllModulByTipe(paramtoken, paramtipe)
	
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrItemNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.Success,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}