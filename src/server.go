package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// 等待组
var waitGroup *sync.WaitGroup

//
// init 初始化函数
//  @Description: 初始化函数
//
func init() {
	waitGroup = new(sync.WaitGroup)
}

//  gin日志结构体
type log struct{}

//
// Write 写入日志数据
//  @Description: 写入日志数据
//  @receiver l 日志结构体实例
//  @param p 写入内容
//  @return n 行数
//  @return err 错误信息
//
func (l *log) Write(p []byte) (n int, err error) {
	str := string(p)
	fmt.Println(strings.ReplaceAll(str, "\n", ""))
	return len(p), nil
}

//
// RunServer 运行服务器
//  @Description: 运行服务器
//
func RunServer(config *Config) {
	waitGroup.Add(1)
	gin.SetMode(config.Server.Mode)
	// 设置默认日志输出为服务器日志输入
	gin.DefaultWriter = new(log)
	// 创建默认gin引擎
	server := gin.Default()
	if config.Server.Static.Open {
		server.Static(config.Server.Static.Path, config.Server.Static.Dir)
	}
	// 设置VNC代理
	server.GET("/vnc", func(ctx *gin.Context) {
		// 此处根据业务对当前用户对应的VNC地址和密码进行映射
		VNCProxyHandle(ctx.Writer, ctx.Request, "10.225.6.253:25071", "spsess@kvm123")
	})
	// 运行服务器
	if e := server.Run(config.Server.Listener); e != nil {
		panic(e)
	}
	waitGroup.Wait()
}
