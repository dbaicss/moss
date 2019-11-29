package services

import (
	"moss-service/db"
	"moss-service/models"
)

func GetDeploymentUpdateHistoryList(name string) ([]*models.UpdateHistory, error)  {
	var dpUpdateHistoryList []*models.UpdateHistory
	dpUpdateHistoryList, err := db.GetDeploymentUpdateHistoryList(name)
	if err != nil {
		return nil, err
	}
	return dpUpdateHistoryList, err
}

func DeploymentUpdateHistory(his *models.UpdateHistory) (err error) {
	 return db.DeploymentUpdateHistory(his)
}