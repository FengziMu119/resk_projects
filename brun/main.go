package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "resk_projects"
	"resk_projects/infra"
	"resk_projects/infra/base"
)

func main() {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
