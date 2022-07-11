package util

import "github.com/bwmarrin/discordgo"

// Commands maps command names to functions
var Commands map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string)

func init() {
	Commands = make(map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string))
}

// IncludesPerm returns whether perm is included in permissions
// discordgo has constants so you can easily fill in perm
func IncludesPerm(perm int, permissions int) bool {
	return (permissions & perm) == perm
}