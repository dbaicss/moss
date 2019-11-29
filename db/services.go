package db

import (
	"moss-service/models"
	"github.com/jinzhu/gorm"
)


//获取所有服务列表
func GetServicesList() ([]*models.Services, error)  {
	var servicesList []*models.Services

	err :=  db.Find(&servicesList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return  servicesList, err
}


//新增deployment记录
func DeploymentCreate(ser *models.Services) (err error) {
	err =  db.FirstOrCreate(ser).Error
	if err != nil  {
		return err
	}
	return err
}

//根据名称获取数据
func GetServiceByName(name string) (*models.Services, error) {
	var services models.Services

	err :=  db.Where("name = ?", name).Find(&services).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &services, err
}