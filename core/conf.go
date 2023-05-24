package core

import (
	_ "errors"
	"fmt"
	"go_vue_blog/config"
	"go_vue_blog/global"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		_ = fmt.Errorf("get yamlConf error: %s", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
	fmt.Println(global.Config)
}

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
