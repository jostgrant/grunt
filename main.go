package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/jostgrant/grunt/cmds"

	"github.com/bwmarrin/discordgo"
	"github.com/jostgrant/grunt/util"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	s, err := discordgo.New("Bot " + util.Config.Token)
	if err != nil {
		log.Panicf("Error creating session: %s", err)
	}

	s.AddHandler(ready)
	s.AddHandler(messageCreate)

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuilds)

	err = s.Open()
	if err != nil {
		log.Panicf("Unable to open session: %s", err)
	}

	// Wait for os terminate events, cleanly close connection when encountered
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	<-closeChan
	log.Print("OS termination received, closing WS and DB")
	s.Close()
	log.Print("Connections closed, bye bye")
}

func ready(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Game: &discordgo.Game{
			Name: "I am a cool bot",
			// Playing 0; Streaming 1; Listening to 2; Custom 4; Competing in 5
			Type: 0,
		},
	})

	log.Print(s.State.User.Username + " is online")
}

var prefix = "bot "
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	perms, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)
	if err != nil {
		log.Printf("Could not get perms for channel %s: %s", m.ChannelID, err)
		return
	}

	perm := util.IncludesPerm

	if perm(discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionEmbedLinks, perms) {
		args := strings.Split(m.Content[len(prefix):], " ")
		if cmd, ok := util.Commands[args[0]]; ok {
			if len(args) == 1 {
				args = []string{}
			} else {
				args = args[1:]
			}

			cmd(s, m, args)
		}
	} else if perm(discordgo.PermissionSendMessages, perms) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I seem to be missing permissions. Below, false indicates a lacking permission:\n```" +
		"Read Text Channels & See Voice Channels: %t\nSend Messages: true\nEmbed Links: %t```",
		perm(discordgo.PermissionViewChannel, perms), perm(discordgo.PermissionEmbedLinks, perms)))
	}
}
