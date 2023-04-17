package commands

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) error {
	help_string := "- Ping : Checks the bot latency"
	_, err := s.ChannelMessageSend(m.ChannelID, help_string)
	return err
}
