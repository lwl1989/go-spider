package config

type LogFileConfig struct {
	FilePath string `json:"file_path,omitempty"` // default /tmp/{YmdHis}.log
}


type LogConfigInterface interface {
	GetLogConfig() (string,interface{})
}

func (file *LogFileConfig) GetLogConfig() (string,interface{}) {
	return "file", file
}

func (file *LogFileConfig) GetFilePath() interface{} {
	return file.FilePath
}
