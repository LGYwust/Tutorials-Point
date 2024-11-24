package flags

import (
	"data_ks2/global"
	"data_ks2/models"
	"fmt"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.ImageModel{},
		&models.MovieModel{},
		&models.ConsultationModel{},
		&models.UserinfoModel{},
	)
	if err != nil {
		fmt.Print("！！数据库迁移失败！！")
	} else {
		fmt.Print("--数据库迁移成功--")
	}

}
