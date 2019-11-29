# kubeconfig

    获取kubernetes配置文件kubeconfig的绝对路径。一般路径为$HOME/.kube/config。该文件主要用来配置本地连接的kubernetes集群。
    
开发环境内容如下：

```
可通过如下命令查看：
kubectl config view
```

```
apiVersion: v1
clusters:
- cluster:
    server: http://k8s-master.dev.icarbonx.net:8080
  name: default
contexts:
- context:
    cluster: default
    namespace: default
    user: ""
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: jane
  user: {}

```

# client-go

client-go是一个调用kubernetes集群资源对象API的客户端，即通过client-go实现对kubernetes集群中资源对象（包括deployment、service、ingress、replicaSet、pod、namespace、node等）的增删改查等操作。大部分对kubernetes进行前置API封装的二次开发都通过client-go这个第三方包来实现。

client-go官方文档：

	https://github.com/kubernetes/client-go
	https://godoc.org/k8s.io/client-go/kubernetes


```
go get k8s.io/client-go@master
```

示例代码：

```
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

```



