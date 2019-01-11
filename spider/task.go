package spider


type Task struct {
	Rule 	*Rule `json:"rule"`
	Uuid    string `json:"uuid"`
	RunTime int64 `json:"run_time"`
	Spacing int64 `json:"spacing_time"`
	EndTime int64 `json:"end_time"`
}