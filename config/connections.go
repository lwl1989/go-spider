package config


type Connections struct {
	Mysql struct{
		Host  string `json:"host,omitempty"`
		Db    string `json:"db,omitempty"`
		User  string `json:"user,omitempty"`
		Pw    string `json:"password,omitempty"`
		TZone string `json:"time_zone,omitempty"`
	} `json:",omitempty"`
	Redis struct{
		Host string `json:"host,omitempty"`
		Db   string `json:"db,omitempty"`
		Pw   string `json:"pw,omitempty"`
	}
}