package main

import (
	"data_ks2/core"
	"data_ks2/flags"
	"data_ks2/global"
	"data_ks2/routers"
)

func main() {
	global.Config = core.InitConfig()
	global.DB = core.InitMySql()
	global.Redis = core.InitRedis()
	routers := routers.Routers()

	opt := flags.Parse()
	if flags.Run(*opt) {
		return
	}
	addr := global.Config.System.Addr()
	routers.Run(addr)
}
