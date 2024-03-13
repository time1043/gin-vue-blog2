package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"time1043/gvb-server/config"
	"time1043/gvb-server/global"
)

const ConfigFile = "settings.yaml"

// 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")

	global.Config = c // 配置文件读取后要写入全局
}

// 修改yaml文件的配置
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}

	global.Log.Info("配置文件修改成功")
	return nil

}
