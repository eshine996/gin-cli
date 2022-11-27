package global

/*
	全局变量存放
*/

var Config *ViperConfig

//进行项目的初始
func init() {
	//初始化配置文件
	if Config == nil {
		Config = NewDefaultConfig()
		if err := Config.ReadData(); err != nil {
			panic(err)
		}
	}
}
