package services

import (
	"moss-service/models"
	"moss-service/db"
)




//获取所有服务列表
func GetServicesList() ([]*models.Services, error)  {
	var servicesList []*models.Services
	servicesList, err := db.GetServicesList()
	if err != nil {
		return nil, err
	}
	return servicesList, err
}

func DeploymentCreate(ser *models.Services) (err error) {
	return db.DeploymentCreate(ser)
}