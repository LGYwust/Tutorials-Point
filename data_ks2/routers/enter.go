package routers

import (
	"data_ks2/api/api_token"
	"data_ks2/routers/image_routers"
	"data_ks2/routers/movie_routers"
	"data_ks2/routers/user_routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Static("/upload", "upload")
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	r.POST("/Checktoken", api_token.CheckToken)
	user_routers.UserRoutersInit(r)
	image_routers.ImageRouterInit(r)
	movie_routers.MovieRouterInit(r)
	return r
}
