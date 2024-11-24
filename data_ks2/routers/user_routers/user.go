package user_routers

import (
	"data_ks2/api/api_user"
	"data_ks2/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	gr := r.Group("User")
	{
		gr.POST("/register", api_user.ApiRegister)
		gr.POST("/login", api_user.ApiLogin)
		gr.POST("/edituser", middleware.JwtAdmin, api_user.ApiEdituser)
		gr.POST("/userinfo", middleware.JwtAdmin, api_user.Saveuserinfo)
		gr.GET("/getinfo", middleware.JwtAdmin, api_user.Getuserinfo)
		gr.GET("/getalluser", api_user.Getalluserinfo)
		gr.POST("/deltuser", middleware.JwtAdmin, api_user.ApiDeluser)
	}
}
