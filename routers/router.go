package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "moss-service/docs"
	"moss-service/routers/api/v1"
)


func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
        //middleware扩展
        apiv1.GET("/services",v1.GetAllServices)
		apiv1.GET("/pods/:namespace/:name",v1.GetPodsByName)
        apiv1.GET("/services/history/:name", v1.GetDeploymentUpdateHistory)
        apiv1.PUT("/services/deployment",v1.UpdateDeployment)
        apiv1.POST("/services/deployment",v1.CreateDeployment)
    }
	return r
}
