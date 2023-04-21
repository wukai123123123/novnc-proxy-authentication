package main

import (
	"encoding/xml"
	"os"
)

//
// Config 配置文件相关结构体
//  @Description: 配置文件相关结构体
//
type Config struct {
	Server ConfigServer `xml:"server"` // 服务器配置
}

//
// ConfigServer 服务器配置
//  @Description: 服务器配置
//
type ConfigServer struct {
	Listener string       `xml:"listener,attr"` // 监听地址和端口
	Mode     string       `xml:"mode,attr"`     // 服务启动模式
	Static   ConfigStatic `xml:"static"`        // 静态资源文件
}

//
// ConfigStatic 静态资源配置
//  @Description: 静态资源配置
//
type ConfigStatic struct {
	Open bool   `xml:"open,attr"` // 是否开启静态资源
	Dir  string `xml:"dir,attr"`  // 物理文件夹路径
	Path string `xml:"path,attr"` // 请求路径
}

//
// AnalysisConfigXMLFile 解析配置文件XML文件
//  @Description: 解析配置文件XML文件
//  @param file 文件地址
//  @return *Config 配置文件对象
//
func AnalysisConfigXMLFile(file string) *Config {
	c := Config{}
	bs, e := os.ReadFile(file)
	if e != nil {
		panic(e)
	}
	if e := xml.Unmarshal(bs, &c); e != nil {
		panic(e)
	}
	return &c
}
