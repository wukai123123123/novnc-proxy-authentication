package main

import (
	"fmt"
	"time"
)

//
// Debug Debug日志
//  @Description: Debug日志
//  @param msg 模板
//  @param args 参数
//
func Debug(msg string, args ...any) {
	console("DEBUG", msg, args...)
}

//
// Info Info日志
//  @Description: Info日志
//  @param msg 模板
//  @param args 参数
//
func Info(msg string, args ...any) {
	console("INFO", msg, args...)
}

//
// Warn Warn日志
//  @Description: Warn日志
//  @param msg 模板
//  @param args 参数
//
func Warn(msg string, args ...any) {
	console("WARN", msg, args...)
}

//
// Error Error日志
//  @Description: Error日志
//  @param msg 模板
//  @param args 参数
//
func Error(msg string, args ...any) {
	console("ERROR", msg, args...)
}

//
// console 日志汇总输出
//  @Description: 日志汇总输出
//  @param t 日志类型
//  @param m 模板
//  @param args 参数
//
func console(t, m string, args ...any) {
	fmt.Printf("[%s] %s | %s\n", t, time.Now().Format("2006/01/02 - 15:04:05"), fmt.Sprintf(m, args...))
}
