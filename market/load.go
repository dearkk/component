package market

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Mysql struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
	PoolSize string `yaml:"pool_size"`
}

type Param struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
}

type Module struct {
	Name    string  `yaml:"name"`
	Enable  bool    `yaml:"enable"`
	Version string  `yaml:"version"`
	Params  []Param `yaml:"params"`
}
type Config struct {
	Version   string   `yaml: "version"`
	IP        string   `yaml: "ip"`
	Port      int      `yaml: "port"`
	Mysql     Mysql    `json:"mysql"`
	Redis     Redis    `json:"redis"`
	ModuleDir string   `yaml: "moduledir"`
	Modules   []Module `yaml: "modules"`
}

type Load interface {
	Start(db *gorm.DB, klog *log.Entry, param *[]Param, addRoute func(module string, routes []Route))
}
