package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/shlex"
)

var botId string

// Runs the bot
func Run() error {
	// Creates a new session
	session, err := discordgo.New("Bot " + TOKEN) // config.go
	if err != nil {
		return err
	}

	bot, err := session.User("@me")
	if err != nil {
		return err
	}

	botId = bot.ID

	// Adds a message handler
	session.AddHandler(messageHandler)
	// Opens the session
	err = session.Open()
	if err != nil {
		return err
	}

	log.Printf("%v is running!\n\n", bot.Username)

	return nil
}

// Handles the messages sent in channel of discord servers that the bot is in it
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, PREFIX) { // config.go
		if m.Author.ID == botId {
			return
		}

		// Divides the arguments removing the prefix "sh!"
		args, err := shlex.Split(m.Content[3:])
		if err != nil {
			return
		}

		// Check if the command required by user is in validCommands
		if command, isValid := validCommands[strings.ToLower(args[0])]; isValid {
			// Removes the command name in args
			args = args[1:]
			err = command(s, m, args)
			if err != nil {
				_ = s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			}

		} else {
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		}
	}
}
