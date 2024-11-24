package image_routers

import (
	"data_ks2/api/api_image"
	"data_ks2/api/api_movies"
	"data_ks2/middleware"

	"github.com/gin-gonic/gin"
)

func ImageRouterInit(r *gin.Engine) {
	gr := r.Group("/image")
	{
		gr.POST("/upload", middleware.JwtAdmin, api_image.ApiIamgeUpload)
		gr.POST("/uploadmovie", api_image.ApiMovieUpload)
		gr.POST("/uploadconsultation", api_image.ApiConsultationUpload)
		gr.POST("/getsearchmovie", api_movies.ApiGetsearchmovies)
		gr.POST("/getselectmovie", api_movies.ApiGetselectmovies)
		gr.GET("/getimage", middleware.JwtAdmin, api_image.ApiGetImage)
		gr.GET("/gethotmovies", api_movies.ApiGethotmovies)
		gr.GET("/getrecommendmovies", api_movies.ApiGetRecommendmovies)
		gr.GET("/getallmovies", api_movies.ApiGetallmovies)
		gr.GET("/getconsultation", api_movies.ApiGetconsult)
	}
}
