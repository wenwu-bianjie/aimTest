package main

import (
	"flag"
	"fmt"
	"github.com/wenwu-bianjie/aimTest/demo1/config"
	"github.com/wenwu-bianjie/aimTest/demo1/dbs"
	//_ "github.com/wenwu-bianjie/aimTest/dome1/dbs"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	flag.StringVar(&confFile, "config", "./config.json", "指定config.json")
	flag.Parse()
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()
	// 加载配置
	if err = config.InitConfig(confFile); err != nil {
		goto ERR
	}

	//连接数据库
	if err = dbs.InitDbs(); err != nil {
		goto ERR
	}

	return
ERR:
	fmt.Println(err)
}
