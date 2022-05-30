package utils

import (
	"go/pkg/mylog"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var log = mylog.NewLog("Warning")

type Conf struct {
	// Mysqlhost   string `yaml:"mysqlhost"`
	// Mysqlpwd    string `yaml:"mysqlpwd"`
	// Mysqldbname string `yaml:"mysqldbname"`
	Mysql	string `yaml:"mysql"`

	Redishost string `yaml:"redishost"`
	Redispwd  string `yaml:"redispwd"`
	Redisdb   int    `yaml:"redisdb"`

	Emailsender    string `yaml:"emailsender"`
	Emailspassword string `yaml:"emailspassword"`
	Emailsmtpaddr  string `yaml:"emailsmtpaddr"`
	Emailsmtport  int    `yaml:"emailsmtport"`
	Port           string    `yaml:"port"`
}

var c Conf

func init() {
	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	//yamlFile, err := ioutil.ReadFile("./../conf.yaml")
	//   yamlFile, err :=  ioutilReadDir("../conf.yaml")
	//  ioutil.ReadAll()
	//yamlFile, err := ioutil.ReadFile("C:\\Users\\goodman\\Desktop\\jsb\\conf.yaml")

	if err != nil {
		log.Error(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetMysql() string {
	return c.Mysql
}

func GetRedishost() string {
	return c.Redishost
}
func GetRedispwd() string {
	return c.Redispwd
}
func GetRedisdb() int {
	return c.Redisdb
}
func GetEmailsender() string {
	return c.Emailsender
}
func GetEmailspassword() string {
	return c.Emailspassword
}
func GetEmailsmtpaddr() string {
	return c.Emailsmtpaddr
}
func GetEmailsmtpport() int {
	return c.Emailsmtport
}
func Getport() string {
	return c.Port
}
