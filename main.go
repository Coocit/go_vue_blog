package main

import (
	"go_vue_blog/core"
	"go_vue_blog/flag"
	"go_vue_blog/global"
	"go_vue_blog/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("go_vue_blog运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
