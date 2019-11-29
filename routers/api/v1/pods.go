package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moss-service/models"
	"moss-service/pkg/app"
	"moss-service/pkg/e"
	"net/http"
)

func GetPodsByName(c *gin.Context)  {
	appG := app.Gin{C: c}
	//获取ns和deployment name参数
	ns := c.Param("namespace")
	name := c.Param("name")
	var (
		podList  []*models.Pods
		err error
	)
	podList, err = G_client.GetPodsByName(ns,name)
	if err != nil {
		fmt.Printf("get all pods failed,err:%v\n",err)
	}
	fmt.Printf("podList from k8s:%#v\n",podList)
	data := make(map[string] interface{})
	data["data"] = podList
	data["code"] = "200"
	data["msg"] = ""
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
