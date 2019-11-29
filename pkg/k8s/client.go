package k8s

import "k8s.io/client-go/kubernetes"

type Client struct {
	K8sClient *kubernetes.Clientset
}
