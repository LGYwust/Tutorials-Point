package service_movie

import (
	"data_ks2/global"
	"data_ks2/models"
	"fmt"
)

func Getconsultation() []models.ConsultationModel {
	var consults []models.ConsultationModel
	err := global.DB.Find(&consults).Error
	if err != nil {
		return nil
	}
	for i := range consults {
		consults[i].Image = fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", consults[i].Image)
	}
	return consults
}
