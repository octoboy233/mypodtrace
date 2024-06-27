package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Exporter struct {
		Jaeger struct {
			Endpoint string `yaml:"endpoint"`
		} `yaml:"jaeger"`
	} `yaml:"exporter"`
}

var SysConfig = new(Config)

func init() {
	// 读取YAML文件内容
	yamlFile, err := ioutil.ReadFile("sysconfig.yaml")
	if err != nil {
		panic(err)
	}

	// 解析YAML内容
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

}
