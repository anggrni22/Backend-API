package router

import (
	"akrab-bangkit2022-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, c controller.UserController) *gin.Engine{
	var gr = r.Group("api")
	{
		gr.POST("register", c.Register)
		gr.POST("login", c.LoginEmail)
	}
	return r
}

func ModulandQuizRouter(r *gin.Engine, c controller.ModulAndQuizController) *gin.Engine{
	var gr = r.Group("api")
	{
		gr.GET("modulQuiz/:level", c.FindAllModulAndQuizByLevel)
		gr.GET("modul", c.FindAllModul)
		gr.GET("modul/:tipe", c.FindAllModulByTipe)
		gr.GET("level", c.FindAllLevel)
		gr.GET("level/:tipe", c.FindAllLevelByTipe)
	}
	return r
}

