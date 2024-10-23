package initialize

func Init() {
	//日志zap初始化
	initLogger()
	//初始化database连接
	initGormDataBaseConnection()
	registerTables()
}
