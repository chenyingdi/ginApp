package ginApp

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AppConfig    AppConfig    `yaml:"app"`
	ServerConfig ServerConfig `yaml:"server"`
	DBConfig     DBConfig     `yaml:"database"`
	RedisConfig  RedisConfig  `yaml:"redis"`
}

func (c *Config) LoadFile(filename string) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return err
	}

	return nil
}

type AppConfig struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	MchID     string `yaml:"mch_id"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

func (s *ServerConfig) ParseUrl() string {
	return s.Host + ":" + s.Port
}

type DBConfig struct {
	Dialect   string `yaml:"dialect"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DBName    string `yaml:"db_name"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parse_time"`
	Loc       string `yaml:"loc"`
}

func (d *DBConfig) ParseUri() string {
	return d.User + ":" + d.Password +
		"@tcp(" + d.Host + ":" + d.Port + ")/" +
		d.DBName + "?charset=" + d.Charset +
		"&parseTime=" + d.ParseTime + "&loc=" + d.Loc
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func (r *RedisConfig) ParseUrl() string {
	return r.Host + ":" + r.Port
}
