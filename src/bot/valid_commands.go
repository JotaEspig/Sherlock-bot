package bot

import (
	"sherlock-bot/src/bot/commands"
	"sherlock-bot/src/bot/commands/admin"

	"github.com/bwmarrin/discordgo"
)

// Map that contains the bot commands
var validCommands = map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string) error{
	"ping": commands.Ping,
	"help": commands.Help,
	// Anime
	"search_anime":   commands.SearchAnime,
	"top_animes":     commands.TopAnimeByScore,
	"popular_animes": commands.TopAnimeByPopularity,
	"get_anime":      commands.GetAnime,
	// Manga
	"search_manga":   commands.SearchManga,
	"top_mangas":     commands.TopMangaByScore,
	"popular_mangas": commands.TopMangaByPopularity,
	"get_manga":      commands.GetManga,
	// Character
	"search_character":   commands.SearchCharacter,
	"popular_characters": commands.TopCharactersByFavourites,
	"get_character":      commands.GetCharacter,
	// Admin
	"logout": admin.Logout,
}
