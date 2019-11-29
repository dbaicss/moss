package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moss-service/models"
	"moss-service/pkg/app"
	"moss-service/pkg/e"
	"moss-service/pkg/k8s"
	"moss-service/pkg/setting"

	//"moss-service/pkg/setting"
	"moss-service/services"
	"net/http"
	"strings"
	"time"
)


var (
	G_client = &k8s.Client{}
)
//获取deployment list
func GetAllServices(c *gin.Context)  {
	appG := app.Gin{C: c}
	var (
		servicesList  []*models.Services
		err error
	)
	/*
	//屏蔽数据库为空的异常
	kubeService := services.Services{}
	servicesList, err := kubeService.GetServicesList()
	fmt.Printf("servicesList from mysq:%#v\n",servicesList)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SERVICES_FAIL, nil)
		return
	}
	 */
	G_client.K8sClient = setting.KubeSetting.KubeClientSet
	//fmt.Printf("client:%#v\n",G_client)
	if len(servicesList) == 0 {
		//servicesList, err = k8s.GetAllService("default")
		servicesList, err = G_client.GetAllService("default")
		if err != nil {
			fmt.Printf("get all service failed,err:%v\n",err)
		}
		fmt.Printf("servicesList from k8s:%#v\n",servicesList)
	}
	data := make(map[string] interface{})
	data["data"] = servicesList
	data["code"] = "200"
	data["msg"] = ""
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//获取deployment update history
func GetDeploymentUpdateHistory(c *gin.Context)  {
	appG := app.Gin{C: c}
	//获取deployment name参数
	//name := c.Param("name")
	var (
		updateHistoryList  []*models.UpdateHistory
		err error
	)
	//updateHistoryList, err = services.GetDeploymentUpdateHistoryList(name)
	if len(updateHistoryList) == 0  {
		updateHistoryList = append(updateHistoryList,&models.UpdateHistory{
			Name:"qc-web",
			Version:"1.15",
			CreatedTime:time.Unix(time.Now().Unix(),0).Format("2006-01-02 03:04:05"),
			Code:"200",
			Msg:"",
		},
		&models.UpdateHistory{
				Name:"qc-web",
				Version:"1.14",
				CreatedTime:time.Unix(time.Now().Unix()-1000,0).Format("2006-01-02 03:04:05"),
		},
		)
	}
	if err != nil {
		fmt.Printf("get update history failed,err:%v\n",err)
	}
	fmt.Printf("updateHistory from k8s:%#v\n",updateHistoryList)
	data := make(map[string] interface{})
	data["data"] = updateHistoryList
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//update deployment
func UpdateDeployment(c *gin.Context)  {
	appG := app.Gin{C: c}
	var (
		err error
	)
	ns := c.PostForm("namespace")
	name := c.PostForm("name")
	image := c.PostForm("image")
	rs := c.PostForm("rs")
	err = G_client.UpdateDeployment(ns,name,image,rs)
	//dev 环境获取version 信息
	vRes := strings.Split(image,"/")
	vLen := len(vRes) - 1
	ver :=  vRes[vLen]
	//idc 环境获取version信息
	/*
	vRes := strings.Split(image,":")[1]
	 */
	data := make(map[string] interface{})
	tm := time.Unix(time.Now().Unix(),0).Format("2006-01-02 03:04:05")
	if err != nil {
		data["data"] = &models.UpdateResult{
			Name:name,
			Code:"500",
			Msg:err.Error(),
			Image: image,
			Version: ver,
			CreatedTime:tm ,
		}
	}
	data["data"] = &models.UpdateResult{
		Name:name,
		Code:"200",
		Msg:"",
		Version: ver,
		Image: image,
		CreatedTime: tm,
	}
	var updateHis = &models.UpdateHistory{
		Name:name,
		Version:ver,
		Image: image,
		CreatedTime:tm,
	}
	err = services.DeploymentUpdateHistory(updateHis)
	if err != nil {
		fmt.Printf("insert deployment update record failed,err:%v\n",err)
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func CreateDeployment(c *gin.Context)  {
	appG := app.Gin{C: c}
	var (
		ret *models.CreateResult
		//dr  *models.Services
		//err error
	)
	yaml := c.PostForm("yaml")
	//fmt.Printf("yaml content:%s\n",yaml)
	ret,_,_ = G_client.CreateDeploy(yaml)
	data := make(map[string] interface{})
	data["data"] = ret
	//更新mysql
	//err = services.DeploymentCreate(dr)
	//if err != nil {
	//	fmt.Printf("insert deployment create record failed,err:%v\n",err)
	//}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
