package config

type LogFileConfig struct {
	FilePath string `json:"file_path,omitempty"` // default /tmp/{YmdHis}.log
}

type LogMgoConfig struct {
	Db   string `json:"db"`
	Collection string `json:"collection"`
}

type LogMysqlConfig struct {
	Connection *Connections
	Table string `json:"table"`
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