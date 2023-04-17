package admin

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Logout(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) error {
	if isOwner(m) {
		_ = s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
		err := s.Close()
		if err != nil {
			_ = s.MessageReactionRemove(m.ChannelID, m.ID, "✅", s.State.User.ID)
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return err
		}

		log.Printf("%v has log out\n", s.State.User.Username)
		os.Exit(0)
	}

	return nil
}

func isOwner(m *discordgo.MessageCreate) bool {
	for _, owner := range Owners.Users {
		if m.Author.ID == owner.Id {
			return true
		}
	}

	return false
}
