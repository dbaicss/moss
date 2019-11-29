package models

type UpdateHistory struct {
	ID 					int `gorm:"primary_key" "json:"id" `
	Name 				string `json:"name"`
	Version				string `json:"status"`
	Image				string `json:"image"`
	CreatedTime			string `json:"created_time"`
	Code string `json:"code"`
	Msg string `json:"msg"`
}

type UpdateResult struct {
	Name 				string `json:"name"`
	Code string `json:"code"`
	Msg string `json:"msg"`
	Version				string `json:"version"`
	Image string `json:"image"`
	CreatedTime			string `json:"created_time"`
}

type CreateResult struct {
	Name 				string `json:"name"`
	Code string `json:"code"`
	Msg string `json:"msg"`
	Image				string `json:"image"`
	CreatedTime			string `json:"created_time"`
}
