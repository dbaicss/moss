package k8s

import (
	"encoding/json"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"moss-service/models"
	"strconv"
	"strings"
	"time"

	"fmt"
	"k8s.io/client-go/util/retry"
	"log"
	"moss-service/pkg/setting"
)


// 获取指定deployment的pod
func (c *Client)GetPodsByName(ns string,name string)(pds []*models.Pods,err error) {
	podList, err := c.K8sClient.CoreV1().Pods(ns).List(metav1.ListOptions{
		LabelSelector:fmt.Sprintf("app=%s",name),
	})
	//podList, err := setting.KubeSetting.KubeClientSet.CoreV1().Pods(ns).List(metav1.ListOptions{})
	fmt.Printf("pod list length:%d\n",len(podList.Items))
	for _,pods := range podList.Items {
		pod := models.Pods{
			Name:pods.Name,
			Status:string(pods.Status.Phase),
			RestartTimes:pods.Status.ContainerStatuses[0].RestartCount,
			CreatedTime:time.Unix(pods.CreationTimestamp.Time.Unix(),0).Format("2006-01-02 03:04:05"),
		}
		pds = append(pds,&pod)
	}
	return
}

//获取所有的ns
func (c *Client)GetAllNamespace()(nsList []string,err error)  {
	fmt.Printf("k8s client set:%#v\n",setting.KubeSetting.KubeClientSet)
	nameSpaceList,err := c.K8sClient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("list deployment failed,err:%v\n",err)
		return
	}
	for _, item := range nameSpaceList.Items {
		//判断非kube开头或者test开头的ns
		if !strings.HasPrefix(item.Name,"kube") || !strings.HasPrefix(item.Name,"test") {
			nsList = append(nsList,item.Name)
		}
	}
	return
}

//获取所有的deployment
func (c *Client)GetAllService(ns string)(serviceList []*models.Services,err error)  {
	svList,err := c.K8sClient.AppsV1().Deployments(ns).List(metav1.ListOptions{})
	for _,sl := range svList.Items {
		svl := models.Services{
			Name:sl.Name,
			Namespace:sl.Namespace,
			SelfLink:sl.SelfLink,
			Replicas:sl.Status.Replicas,
			Image:sl.Spec.Template.Spec.Containers[0].Image,
			CreatedTime:time.Unix(sl.ObjectMeta.CreationTimestamp.Time.Unix(),0).Format("2006-01-02 03:04:05"),
			UpdateTime:time.Unix(sl.Status.Conditions[0].LastUpdateTime.Time.Unix(),0).Format("2006-01-02 03:04:05"),
		}
		serviceList = append(serviceList,&svl)
	}
	if err != nil {
		log.Printf("list deployment failed,err:%v\n",err)
		return
	}
	return
}

//更新deployment
func (c *Client)UpdateDeployment(ns,name,image,replicatset string) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr :=c.K8sClient.AppsV1().Deployments(ns).Get(name, metav1.GetOptions{})
		//deployment不存在或者获取失败
		if getErr != nil || errors.IsNotFound(getErr) {
			fmt.Errorf("Failed to get latest version of Deployment or Deployment does not exist: %v", getErr)
			return  getErr
		}
		if len(replicatset) == 0 {
			result.Spec.Replicas = result.Spec.Replicas
		}
		rs,err := strconv.ParseInt(replicatset,10,32)
		if err != nil {
			result.Spec.Replicas = result.Spec.Replicas
		}
		//更改实例数和镜像名称
		result.Spec.Replicas = int32Ptr(int32(rs))                           // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = image         // change image version
		_, updateErr := setting.KubeSetting.KubeClientSet.AppsV1().Deployments(ns).Update(result)
		return updateErr
	})
	if err != nil {
		fmt.Printf("Update deployment %s failed: %v",name, err)
		return
	}
	fmt.Printf("Updated deployment %s succ...",name)
	return
}

// add deployment
func (c *Client)CreateDeploy(yamlContent string)(*models.CreateResult,*models.Services, error) {
	var (
		cr  = models.CreateResult{}
		dr = models.Services{}
		deployJson []byte
		deployment  = appsv1.Deployment{}
		err error
	)
	//yaml转json
	if deployJson, err = yaml2.ToJSON([]byte(yamlContent)); err != nil {
		cr.Code = "500"
		cr.Msg = "yaml to json failed"
		fmt.Printf("yaml to json failed,err:%v\n",err)
		return &cr,nil,err
	}

	//json转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		cr.Code = "500"
		cr.Msg = "json to struct failed"
		fmt.Printf("json to struct failed,err:%v\n",err)
		return &cr,nil,err
	}
	//fmt.Printf("deployment result:%#v\n",deployment)
	tm := time.Unix(time.Now().Unix(),0).Format("2006-01-02 03:04:05")
	//fmt.Printf("deployment name:%s,deployment namespace:%s\n",deployment.ObjectMeta.Name,deployment.ObjectMeta.Namespace)
	if _, err = c.K8sClient.AppsV1().Deployments(deployment.Namespace).Get(deployment.Name, metav1.GetOptions{}); err != nil {
		//deployment已存在
		if !errors.IsNotFound(err) {
			cr.Name = deployment.Name
			cr.Image = deployment.Spec.Template.Spec.Containers[0].Image
			cr.Code = "404"
			cr.Msg = fmt.Sprintf("deployment %s does not exist",deployment.Name)
			cr.CreatedTime = tm
			fmt.Printf("deployment %s is exist,create faild,err:%v\n",deployment.Name,err)
			return &cr,nil,err
		}
		//deployment不存在则创建
		//fmt.Printf("before deployment created:%#v\n",deployment)
		if _, err = c.K8sClient.AppsV1().Deployments(deployment.Namespace).Create(&deployment); err != nil {
			cr.Name = deployment.Name
			cr.Image = deployment.Spec.Template.Spec.Containers[0].Image
			fmt.Printf("deployment image:%s\n",deployment.Spec.Template.Spec.Containers[0].Image)
			cr.Code = "500"
			cr.Msg = fmt.Sprintf("deployment %s create failed,err:%v",deployment.Name,err)
			cr.CreatedTime = tm
			fmt.Printf("deployment %s create faild,err:%v\n",deployment.Name,err)
			return &cr,nil,err
		}
	}
	//fmt.Printf("after deployment created:%#v\n",deployment)
	cr.Name = deployment.Name
	cr.Image = deployment.Spec.Template.Spec.Containers[0].Image
	cr.Code = "200"
	cr.Msg = fmt.Sprintf("deployment %s create succ",deployment.Name)
	cr.CreatedTime = tm
	dr.Name = deployment.Name
	dr.Namespace = deployment.Namespace
	dr.SelfLink = deployment.ObjectMeta.SelfLink
	dr.Replicas = *deployment.Spec.Replicas
	dr.Image = deployment.Spec.Template.Spec.Containers[0].Image
	dr.CreatedTime = tm
	dr.UpdateTime = tm
	fmt.Printf("deployment %s create succ...",deployment.Name)
	return &cr,&dr,nil
}

func int32Ptr(i int32) *int32 { return &i }
