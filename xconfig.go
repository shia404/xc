package xc

type XConfig struct {
	Env  string `json:"Env,default=main"`
	Gorm GormConfig
}
