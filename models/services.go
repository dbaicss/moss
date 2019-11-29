package models

/*
name:K8S服务
*/
type Services struct {
	ID 					int `gorm:"primary_key" "json:"id" `
	Name 				string `json:"name"`
	Namespace        	string `json:"namespace"`
	SelfLink 		string `json:"description"`
	Replicas			int32 `json:"replicas"`
	Image				string `json:"image"`
	CreatedTime			string `json:"created_time"`
	UpdateTime			string `json:"update_time"`
}



type GetServices struct {
	Name 				string `json:"name"`
	Namespace        	string `json:"namespace"`
	SelfLink 		string `json:"description"`
	Replicas			int32 `json:"replicas"`
	Image				string `json:"image"`
	CreatedTime			string `json:"created_time"`
	UpdateTime			string `json:"update_time"`
	Code   string `json:"code"`
	Msg  string `json:"msg"`
}