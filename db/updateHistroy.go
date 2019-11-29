package db

import (
	"github.com/jinzhu/gorm"
	"moss-service/models"
)

//获取deployment更新记录
func GetDeploymentUpdateHistoryList(name string) (history []*models.UpdateHistory, err error) {
	err =  db.Where("name = ?", name).Order("created_time desc").Limit(5).Find(history).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return
}

//新增deployment更新记录
func DeploymentUpdateHistory(history *models.UpdateHistory) (err error) {
	err =  db.FirstOrCreate(history).Error
	if err != nil  {
		return err
	}
	return err
}