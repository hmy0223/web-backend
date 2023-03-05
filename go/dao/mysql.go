package dao

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v3"
)

type config struct {
	DbName     string `yaml:"dbName"`
	DbPassword string `yaml:"dbPassword"`
	DbHost     string `yaml:"dbHost"`
	DbPort     string `yaml:"dbPort"`
	UserName   string `yaml:"userName"`
}

const DRIVER = "mysql"

var SqlSession *gorm.DB

func (c *config) getConfig() *config {
	yamlFile, err := os.ReadFile("resources/application.yaml")

	if err != nil {
		fmt.Println(err.Error())
	}

	yamlError := yaml.Unmarshal(yamlFile, c)

	if yamlError != nil {
		fmt.Println(yamlError.Error())
	}
	return c

}

func InitMySql() (err error) {
	var c config
	conf := c.getConfig()
	// map application yaml
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)
	// dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)

	// connect to mySQL
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	return SqlSession.DB().Ping()
}

func Close() {
	SqlSession.Close()
}
