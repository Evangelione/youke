package infra

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App   AppConf
	Mysql MysqlConf
	Redis RedisConf
}

type AppConf struct {
	Name string
	Port string
}

type MysqlConf struct {
	Address  string
	UserName string
	Password string
	DataBase string
}

type RedisConf struct {
	Address string
}

var Conf *Config

func LoadConfig(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = yaml.Unmarshal(yamlFile, &Conf); err != nil {
		log.Fatal(err.Error())
	}
	return Conf
}
