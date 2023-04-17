package commands

import (
	"errors"
	"fmt"
	"sherlock-bot/src/anilistapi"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SearchCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	var (
		page      int
		perPage   int
		strToSend string
		err       error
	)

	if len(args) > 3 {
		return errors.New("params: Too many arguments in the command")
	}
	if len(args) < 1 {
		return errors.New("params: Missing arguments")
	}

	search := args[0]
	if len(args) >= 2 {
		page, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		if len(args) >= 3 {
			perPage, err = strconv.Atoi(args[2])
			if err != nil {
				return err
			}
		}
	}

	response, err := anilistapi.SearchCharacter(search, page, perPage)
	if err != nil {
		return err
	}

	characters := response.Page.Characters

	// Creates the string to send in the channel
	for _, character := range characters {
		strToSend += fmt.Sprintf("%v - ( %v )\n",
			character.Id, character.Name.Full)
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

func TopCharactersByFavourites(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	var (
		page      int
		perPage   int
		strToSend string
		err       error
	)

	if len(args) > 2 {
		return errors.New("params: Too many arguments in the command")
	}

	if len(args) >= 1 {
		page, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if len(args) == 2 {
			perPage, err = strconv.Atoi(args[1])
			if err != nil {
				return err
			}
		}
	}

	response, err := anilistapi.TopCharactersByFavourites(page, perPage)
	if err != nil {
		return err
	}

	characters := response.Page.Characters

	// Creates the string to send in the channel
	for _, character := range characters {
		strToSend += fmt.Sprintf("%v - ( %v | %v)\n",
			character.Id, character.Name.Full, character.Favourites)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

func GetCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	var mediaToString string

	characterid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	response, err := anilistapi.GetCharacter(characterid)
	if err != nil {
		return err
	}

	character := response.Character

	// Creates embed message variables
	var embedVar discordgo.MessageEmbed

	// Title
	embedVar.Title = character.Name.Full

	// Description
	anilistapi.TreatDescription(character.Description, &embedVar.Description)

	// Color
	embedVar.Color = 0x222267

	// Fields
	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Native Name", Value: character.Name.Native, Inline: false})

	if character.Age != "" {
		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "Age", Value: character.Age, Inline: false})
	}

	if character.DateOfBirth.Day != 0 && character.DateOfBirth.Month != 0 {
		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "Date of birth", Value: fmt.Sprintf(
				"%v %v",
				time.Month(character.DateOfBirth.Month).String(), character.DateOfBirth.Day,
			), Inline: true})
	}

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Gender", Value: character.Gender, Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Favourites", Value: fmt.Sprint(character.Favourites)})

	// Media
	for idx, animanga := range character.Media.Nodes {
		if idx < 5 {
			mediaToString += fmt.Sprintf("%v - ( %v | %v)\n",
				animanga.Id, animanga.Title.Romaji, animanga.Format)
		} else {
			mediaToString += "And more..."
			break
		}
	}
	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Media", Value: mediaToString})

	// Image
	embedVar.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: character.Image.Large}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embedVar)
	if err != nil {
		return err
	}

	return nil

}
