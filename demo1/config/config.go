package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MysqlUserName string `json:"mysql_user_name"`
	MysqlPassWord string `json:"mysql_pass_word"`
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     string `json:"mysql_port"`
	MysqlDatabase string `json:"mysql_database"`
	MysqlCharset  string `json:"mysql_charset"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 1, 把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2, 做JSON反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3, 赋值单例
	G_config = &conf

	return
}
