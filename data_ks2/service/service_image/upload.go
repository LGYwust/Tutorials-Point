package service_image

import (
	"data_ks2/global"
	"data_ks2/models"
	"mime/multipart"
)

func CheckImage(filehash string) string {
	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, "hash = ?", filehash).Error
	if err == nil {
		return ""
	}
	return imageModel.Path
}
func SaveImage(userID float64, file *multipart.FileHeader, savePath string, hash string) bool {
	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, "userID = ?", int(userID)).Error
	if err != nil {
		imageModel = models.ImageModel{
			UserID:   int(userID),
			FileName: file.Filename,
			Size:     file.Size,
			Path:     savePath,
			Hash:     hash,
		}
		err = global.DB.Create(&imageModel).Error
		if err == nil {
			return false
		}
		return true
	}
	err = global.DB.Model(&imageModel).Updates(models.ImageModel{
		UserID:   int(userID),
		FileName: file.Filename,
		Size:     file.Size,
		Path:     savePath,
		Hash:     hash,
	}).Error
	if err == nil {
		return false
	}
	return true
}

func SaveMovieImage(movieName string, savePath string) bool {
	var movieModel models.MovieModel
	err := global.DB.Take(&movieModel, "name = ?", movieName).Error
	if err != nil {
		return false
	}
	err = global.DB.Model(&movieModel).Updates(models.MovieModel{
		Image: savePath,
	}).Error
	if err == nil {
		return false
	}
	return true
}

func SaveConsultationImage(title string, savePath string) bool {
	var mc models.ConsultationModel
	err := global.DB.Take(&mc, "title = ?", title).Error
	if err != nil {
		return false
	}
	err = global.DB.Model(&mc).Updates(models.ConsultationModel{
		Image: savePath,
	}).Error
	if err == nil {
		return false
	}
	return true
}
