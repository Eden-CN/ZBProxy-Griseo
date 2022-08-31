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
			{
				Color: chat.Gray,
				Text: fmt.Sprintf("当前时间戳: %d | 玩家名: %s | 服务器: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
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
			{
				Color: chat.Gray,
				Text: fmt.Sprintf("当前时间戳: %d | 玩家名: %s | 服务器: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
		},
	}
}
