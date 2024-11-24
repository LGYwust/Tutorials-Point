package service_image

import (
	"data_ks2/global"
	"data_ks2/models"
	"fmt"
)

func GetImage(userID float64) string {
	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, "userID = ?", int(userID)).Error
	if err != nil {
		return "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
	}
	return fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", imageModel.Path)
}
