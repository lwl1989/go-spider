package config

type LogFileConfig struct {
	FilePath string `json:"file_path,omitempty"` // default /tmp/{YmdHis}.log
}

type LogMgoConfig struct {
	Host string `json:"host,omitempty"`
	Db   string `json:"db,omitempty"`
	Pw   string `json:"pw,omitempty"`
}

type LogMysqlConfig struct {
	Connection *Connections
	Table string `json:"table,omitempty"`
	TaskTable string `json:"task_table,omitempty"`
}

type LogConfigInterface interface {
	GetLogConfig() (string,interface{})
}

func (file *LogFileConfig) GetLogConfig() (string,interface{}) {
	return "file", file
}

func (mysql *LogMysqlConfig) GetLogConfig() (string,interface{}) {
	return "file", mysql
}

func (mongo *LogMgoConfig) GetLogConfig() (string,interface{}) {
	return "file", mongo
}