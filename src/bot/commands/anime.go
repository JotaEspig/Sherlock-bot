package commands

import (
	"errors"
	"fmt"
	"sherlock-bot/src/anilistapi"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func SearchAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.SearchAnime(search, page, perPage)
	if err != nil {
		return err
	}

	animes := response.Page.Media

	// Creates the string to send in the channel
	for _, anime := range animes {
		strToSend += fmt.Sprintf("%v - ( %v | %v )\n",
			anime.Id, anime.Title.Romaji, anime.Format)
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

// Send a message in a discord channel with the animes with the highest scores in Anilist
func TopAnimeByScore(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.TopAnimeByScore(page, perPage)
	if err != nil {
		return err
	}

	animes := response.Page.Media

	// Creates the string to send in the channel
	for _, anime := range animes {
		strToSend += fmt.Sprintf("%v - ( %v | %v | %v%% )\n",
			anime.Id, anime.Title.Romaji, anime.Format, anime.AverageScore)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

// Send a message in a discord channel with the most popular animes in Anilist
func TopAnimeByPopularity(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.TopAnimeByPopularity(page, perPage)
	if err != nil {
		return err
	}

	animes := response.Page.Media

	// Creates the string to send in the channel
	for _, anime := range animes {
		strToSend += fmt.Sprintf("%v - ( %v | %v | %v watched)\n",
			anime.Id, anime.Title.Romaji, anime.Format, anime.Popularity)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

// Send a message in a discord channel with the anime request by the user (using anilist id)
func GetAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	animeid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	response, err := anilistapi.GetAnime(animeid)
	if err != nil {
		return err
	}

	anime := response.Media

	// Creates embed message variables
	var embedVar discordgo.MessageEmbed

	// Title
	embedVar.Title = anime.Title.Romaji

	// Description
	anilistapi.TreatDescription(anime.Description, &embedVar.Description)

	// Color
	embedVar.Color = 0x222267

	// Fields
	if anime.Title.English != "" {
		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "English Name", Value: anime.Title.English})
	}
	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Native Name", Value: anime.Title.Native})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Format", Value: anime.Format, Inline: true})

	if anime.Format != "MOVIE" {
		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "Episodes", Value: fmt.Sprint(anime.Episodes), Inline: true})
	}

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Status", Value: anime.Status, Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Season", Value: anime.Season, Inline: true})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Year", Value: fmt.Sprint(anime.SeasonYear), Inline: true})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Average Score", Value: fmt.Sprintf("%v%%", anime.AverageScore), Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Popularity", Value: fmt.Sprint(anime.Popularity), Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Favourites", Value: fmt.Sprint(anime.Favourites), Inline: true})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Source", Value: anime.Source, Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Genres", Value: fmt.Sprint(anime.Genres), Inline: true})

	// Image
	embedVar.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: anime.CoverImage.ExtraLarge}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embedVar)
	if err != nil {
		return err
	}

	return nil
}
