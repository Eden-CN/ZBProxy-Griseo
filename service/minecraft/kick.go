package minecraft

import (
	"fmt"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/net/packet"
	"github.com/LittleGriseo/GriseoProxy/config"
	"time"
)

func generateKickMessage(s *config.ConfigProxyService, name packet.String) chat.Message {
	return chat.Message{
		Color: chat.White,
		Extra: []chat.Message{
			{Bold: false, Color: chat.Blue, Text: "Griseo"},
			{Bold: false, Text: "Proxy"},
			{Text: " - "},
			{Bold: false, Text: "连接被拒绝！\n"},

			{Text: "您并没有获得此次测试的资格！\n"},
		},
	}
}

func generatePlayerNumberLimitExceededMessage(s *config.ConfigProxyService, name packet.String) chat.Message {
	return chat.Message{
		Color: chat.White,
		Extra: []chat.Message{
			{Bold: false, Color: chat.Blue, Text: "Griseo"},
			{Bold: false, Text: "Proxy"},
			{Text: " - "},
			{Bold: false, Text: "连接被拒绝！\n"},

			{Text: "当前服务器已满人！\n"},
		},
	}
}
