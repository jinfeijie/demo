package dto

type FrameworkConfig struct {
	Server Server               `yaml:"server"`
	Mysql  map[string]mysqlConf `yaml:"mysql"`
}

type Server struct {
	Port   uint64 `yaml:"port"`
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
}

type mysqlConf struct {
	Host     string `yaml:"host"`
	Port     uint64 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}
