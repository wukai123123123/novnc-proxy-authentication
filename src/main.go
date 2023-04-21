package main

//
// main 程序入口函数
//  @Description: 程序入口函数
//
func main() {
	// 读取配置文件
	config := AnalysisConfigXMLFile("./conf/config.xml")
	// 运行服务器
	RunServer(config)
}
