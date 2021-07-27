package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string = os.Getenv("DISCORD_BOT_TOKEN")
)

func main() {
	// Check if token is empty
	err := checkToken()
	if err != nil {
		fmt.Println(err)
	}

	// Create a new discordgo session
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println(err)
	}

	// Add a message handler
	dg.AddHandler(messageCreate)

	// Set intents and establish connection to Discord
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = dg.Open()
	if err != nil {
		fmt.Println(err)
	}

	// Catch signal [SIGINT, SIGTERM] to shut down the process
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close connection to Discord
	defer func() {
		dg.Close()
	}()
}

func checkToken() error {
	if botToken == "" {
		return errors.New("discord bot token is required")
	}

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var prefix string = "$"
	m.Content = strings.ToLower(m.Content)
	if strings.HasPrefix(m.Content, prefix) {
		m.Content = strings.TrimPrefix(m.Content, prefix)
	} else {
		return
	}

	cmdAndArgs := strings.Split(m.Content, " ")
	if cmdAndArgs[0] == "ping" {
		s.ChannelMessageSend(m.ChannelID, "<@!"+m.Author.ID+"> "+"Pong!")
	}
}
