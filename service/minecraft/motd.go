package minecraft

import (
	"encoding/json"
	"github.com/Tnze/go-mc/net/packet"
	"github.com/LittleGriseo/GriseoProxy/config"
	"github.com/LittleGriseo/GriseoProxy/service/transfer"
	"github.com/LittleGriseo/GriseoProxy/version"
)

const DefaultMotd = `data:image/jpeg;base64,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`

type motdObject struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		/*		Sample []struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"sample"`*/
	} `json:"players"`
	Description struct {
		Text string `json:"text"`
	} `json:"description"`
	Favicon string `json:"favicon"`
}

func generateMotdPacket(protocolVersion int, s *config.ConfigProxyService, options *transfer.Options) packet.Packet {
	online := s.Minecraft.OnlineCount.Online
	if online < 0 {
		online = options.GetCount()
	}
	motd, _ := json.Marshal(motdObject{
		Version: struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		}{
			Name:     "ZBProxy " + version.Version,
			Protocol: protocolVersion,
		},
		Players: struct {
			Max    int `json:"max"`
			Online int `json:"online"`
		}{
			Max:    s.Minecraft.OnlineCount.Max,
			Online: int(online),
		},
		Description: struct {
			Text string `json:"text"`
		}{
			Text: s.Minecraft.MotdDescription,
		},
		Favicon: s.Minecraft.MotdFavicon,
	})
	return packet.Marshal(
		0x00, // Client bound : Status Response
		packet.String(motd),
	)
}
