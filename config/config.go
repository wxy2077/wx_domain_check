/*
* @Time    : 2020-08-20 23:30
* @Author  : CoderCharm
* @File    : config.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 读取配置文件
**/
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 定义一个 config结构体变量
var GlobConfig Config

// 声明 配置结构体
type Config struct {
	AppId  string   `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}


// 读取配置文件
func (c *Config) getConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

// 初始化读取配置
func init() {
	// 直接赋值给结构体
	GlobConfig.getConfig()
}
