package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wenwu-bianjie/aimTest/demo1/config"
	"log"
	"time"
)

var MysqlDb *sql.DB

// 初始化链接
func InitDbs() (err error) {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", config.G_config.MysqlUserName,
		config.G_config.MysqlPassWord, config.G_config.MysqlHost,
		config.G_config.MysqlPort, config.G_config.MysqlDatabase, config.G_config.MysqlCharset)

	// 打开连接失败
	MysqlDb, err = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if err != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + err.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if err = MysqlDb.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}
	return
}
