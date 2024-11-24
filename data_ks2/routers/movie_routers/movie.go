package movie_routers

import (
	"data_ks2/api/api_movies"

	"github.com/gin-gonic/gin"
)

func MovieRouterInit(r *gin.Engine) {
	gr := r.Group("/movie")
	{
		gr.POST("/addmovie", api_movies.ApiAddmovie)
		gr.POST("/deltmovie", api_movies.ApiDelmovie)
		gr.POST("/addconsultation", api_movies.ApiAddconsultation)
		gr.POST("/deltconsultation", api_movies.ApiDelconsultation)
	}
}
