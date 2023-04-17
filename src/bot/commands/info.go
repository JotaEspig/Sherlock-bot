package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) error {
	latency := s.HeartbeatLatency().Milliseconds()
	_, err := s.ChannelMessageSend(m.ChannelID, strconv.FormatInt(latency, 10)+"ms")
	return err
}

func About(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) error {
	about_string := "" // TODO Create a string about the bot
	_, err := s.ChannelMessageSend(m.ChannelID, about_string)
	return err
}
