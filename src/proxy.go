package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

//
// VNCProxyHandle VNC代理处理器
//  @Description: VNC代理处理器
//  @param w writer
//  @param r request
//  @param addr VNC TCP地址端口
//  @param password VNC 密码
//
func VNCProxyHandle(w http.ResponseWriter, r *http.Request, addr, password string) {
	ug := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := ug.Upgrade(w, r, nil)
	if err != nil {
		Error("failed to upgrade to WS: %s", err)
		return
	}
	vnc, err := net.Dial("tcp", addr)
	if err != nil {
		Error("failed to bind to the VNC Server: %s", err)
	}
	flag := 0            // 公共标记,用于标记当前执行的步骤 0其他(待机标记) 1握手 2加密方式传递 3加密方式应答 4扰码传递 5密码传递
	var challenge []byte // 密码加密盐
	u, _ := uuid.NewV4() // 生成UUID用于标识当前代理请求
	// 创建协程处理websocket请求
	go func(wsConn *websocket.Conn, conn net.Conn) {
		defer func() {
			// 关闭资源
			if err := recover(); err != nil {
				Error("reading from WS failed: %s", err)
			}
			if conn != nil {
				_ = conn.Close()
			}
			if wsConn != nil {
				_ = wsConn.Close()
			}
		}()
		for {
			if (conn == nil) || (wsConn == nil) {
				return
			}
			if _, buffer, e := wsConn.ReadMessage(); e == nil {
				// 将buffer转为16进制并对比加密格式,TCP返回可选加密格式01,02;客户端需选择02
				if len(password) > 0 && fmt.Sprintf("%x", buffer) == "02" {
					// 对公共标记进行修改
					flag = 3
				}
				// 如果公共标记为4扰码传递,则收到的信息为密码信息
				if flag == 4 && len(challenge) > 0 {
					// 设置标记位
					flag = 5
					// 将密码注入替换为真实密码
					bufferTmp := GetVNCAuthenticationBytes([]byte(password), challenge)
					Debug("%s proxy injection password: %x -> %x", u, buffer, bufferTmp)
					buffer = bufferTmp
				}
				// 转发
				if _, err := conn.Write(buffer); err != nil {
					Error("%s writing to TCP failed: %s", u, err)
				}
			}
		}
	}(ws, vnc)
	// 创建协程处理tcp响应
	go func(wsConn *websocket.Conn, conn net.Conn) {
		var tcpBuffer [2048]byte
		defer func() {
			if conn != nil {
				_ = conn.Close()
			}
			if wsConn != nil {
				_ = wsConn.Close()
			}
		}()
		for {
			if (conn == nil) || (wsConn == nil) {
				return
			}
			n, err := conn.Read(tcpBuffer[0:])
			if err != nil {
				Error("%s reading from TCP failed: %s", u, err)
				return
			}
			// 根据flag状态及响应内容进行处理
			hexBuffer := fmt.Sprintf("%x", tcpBuffer[0:n])
			if flag == 0 && hexBuffer == "0102" {
				flag = 2
			}
			// 接收到密码加密盐
			if flag == 3 && len(hexBuffer) > 0 {
				challenge = tcpBuffer[0:n]
				flag = 4
			}
			// 转发
			if err := wsConn.WriteMessage(websocket.BinaryMessage, tcpBuffer[0:n]); err != nil {
				Error("%s writing to WS failed: %s", u, err)
			}
		}
	}(ws, vnc)
}
