package service

import (
	"fmt"
	"github.com/LittleGriseo/GriseoProxy/config"
	"github.com/LittleGriseo/GriseoProxy/service/minecraft"
	"github.com/LittleGriseo/GriseoProxy/service/tls"
	"github.com/LittleGriseo/GriseoProxy/service/transfer"
	"log"
	"net"
)

func newConnReceiver(s *config.ConfigProxyService,
	conn *net.TCPConn,
	options *transfer.Options) {

	log.Println("服务", s.Name, ": 一个新的连接请求由", conn.RemoteAddr().String(), "接受。")
	defer log.Println("服务", s.Name, ": 一个连接", conn.RemoteAddr().String(), "关闭.")
	var err error // in order to avoid scoop problems
	var remote net.Conn = nil

	if options.IsTLSHandleNeeded {
		remote, err = tls.NewConnHandler(s, conn, options.Out)
		if err != nil {
			conn.Close()
			return
		}
	} else if options.IsMinecraftHandleNeeded {
		remote, err = minecraft.NewConnHandler(s, conn, options)
		if err != nil {
			conn.Close()
			return
		}
	}

	if remote == nil {
		remote, err = options.Out.Dial("tcp", fmt.Sprintf("%v:%v", s.TargetAddress, s.TargetPort))
		if err != nil {
			log.Printf("服务 %s: 拨号到目标服务器失败: %v", s.Name, err.Error())
			conn.Close()
			return
		}
	}
	options.AddCount(1)
	defer options.AddCount(-1)
	transfer.SimpleTransfer(conn, remote, options.FlowType)
}
