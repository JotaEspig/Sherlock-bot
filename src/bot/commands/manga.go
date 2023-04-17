package commands

import (
	"errors"
	"fmt"
	"sherlock-bot/src/anilistapi"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SearchManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.SearchManga(search, page, perPage)
	if err != nil {
		return err
	}

	mangas := response.Page.Media
	// Creates the string to send in the channel
	for _, manga := range mangas {
		strToSend += fmt.Sprintf("%v - ( %v | %v )\n",
			manga.Id, manga.Title.Romaji, manga.Format)
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

func TopMangaByScore(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.TopMangaByScore(page, perPage)
	if err != nil {
		return err
	}

	mangas := response.Page.Media

	// Creates the string to send in the channel
	for _, manga := range mangas {
		strToSend += fmt.Sprintf("%v - ( %v | %v | %v%% )\n",
			manga.Id, manga.Title.Romaji, manga.Format, manga.AverageScore)
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

func TopMangaByPopularity(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
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

	response, err := anilistapi.TopMangaByPopularity(page, perPage)
	if err != nil {
		return err
	}

	mangas := response.Page.Media
	// Creates the string to send in the channel
	for _, manga := range mangas {
		strToSend += fmt.Sprintf("%v - ( %v | %v | %v watched)\n",
			manga.Id, manga.Title.Romaji, manga.Format, manga.Popularity)
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, strToSend)
	if err != nil {
		return err
	}

	return nil
}

// Send a message in a discord channel with the manga request by the user (using anilist id)
func GetManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	mangaid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	response, err := anilistapi.GetManga(mangaid)
	if err != nil {
		return err
	}

	manga := response.Media

	// Creates embed message variables
	var embedVar discordgo.MessageEmbed

	// Title
	embedVar.Title = manga.Title.Romaji

	// Description
	anilistapi.TreatDescription(manga.Description, &embedVar.Description)

	// Color
	embedVar.Color = 0x222267

	// Fields
	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "English Name", Value: manga.Title.English})
	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Native Name", Value: manga.Title.Native})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Format", Value: manga.Format, Inline: true})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Status", Value: manga.Status, Inline: false})

	if manga.Status != "RELEASING" {
		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "Chapters", Value: fmt.Sprint(manga.Chapters), Inline: true})

		embedVar.Fields = append(embedVar.Fields,
			&discordgo.MessageEmbedField{Name: "Volumes", Value: fmt.Sprint(manga.Volumes), Inline: true})
	}

	embedVar.Fields = append(
		embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Start Date",
			Value: fmt.Sprintf("%v %v, %v",
				time.Month(manga.StartDate.Month).String(), manga.StartDate.Day, manga.StartDate.Year), Inline: true},
	)

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Average Score", Value: fmt.Sprintf("%v%%", manga.AverageScore), Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Popularity", Value: fmt.Sprint(manga.Popularity), Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Favourites", Value: fmt.Sprint(manga.Favourites), Inline: true})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Source", Value: manga.Source, Inline: false})

	embedVar.Fields = append(embedVar.Fields,
		&discordgo.MessageEmbedField{Name: "Genres", Value: fmt.Sprint(manga.Genres), Inline: true})

	// Image
	embedVar.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: manga.CoverImage.ExtraLarge}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embedVar)
	if err != nil {
		return err
	}

	return nil
}
