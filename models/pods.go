package models

type Pods struct {
	Name 				string `json:"name"`
	Status				string `json:"status"`
	RestartTimes        int32  `json:"restart_times"`
	CreatedTime			string `json:"created_time"`
}


type GetPods struct {
	Name 				string `json:"name"`
	Status				string `json:"status"`
	RestartTimes        int32  `json:"restart_times"`
	CreatedTime			string `json:"created_time"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
}