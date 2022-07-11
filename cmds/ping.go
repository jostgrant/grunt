package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jostgrant/grunt/util"
)

func init() {
	util.Commands["ping"] = ping
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "Pong...")
	if err != nil {
		return
	}
	msge, err := s.ChannelMessageEdit(m.ChannelID, msg.ID, "Pong?")
	if err != nil {
		return
	}
	
	start, _ := m.Timestamp.Parse()
	end, _ := msg.Timestamp.Parse()
	sendLatency := end.Sub(start)

	start = end
	end, _ = msge.EditedTimestamp.Parse()
	editLatency := end.Sub(start)

	_, err = s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("Pong!\n``` - Send latency: %dms (%dμs)\n - Edit latency: %dms (%dμs)\n - API latency: %dms (%dμs)```",
	sendLatency / 1e6, sendLatency / 1e3, editLatency / 1e6, editLatency / 1e3, s.HeartbeatLatency().Milliseconds(), s.HeartbeatLatency().Microseconds()))
}
