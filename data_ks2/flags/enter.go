package flags

import "flag"

type Option struct {
	DB   bool
	Port int
	Load string
}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.IntVar(&option.Port, "port", 0, "设定端口")
	flag.StringVar(&option.Load, "load", "", "导入SQL数据库")
	flag.Parse()
	return option
}

func Run(option Option) bool {
	if option.DB {
		DB()
		return true
	}
	return false
}
