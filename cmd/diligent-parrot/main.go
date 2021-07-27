package main

import (
	"errors"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rti56kt/diligent-parrot/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken   string = os.Getenv("DISCORD_BOT_TOKEN")
	logVerbose bool
)

func init() {
	flag.BoolVar(&logVerbose, "l", false, "Log level (false:INFO true:DEBUG)")
	flag.Parse()

	loggerSet()
	logger.Logger.WithField("type", "process").Info("diligent-parrot start, hi, let's party!")
	logger.Logger.WithField("type", "flag").Debug("logVerbose is set to: ", logVerbose)
}

func loggerSet() {
	if logVerbose {
		logger.Logger.SetLevel(logger.Debug)
	} else {
		logger.Logger.SetLevel(logger.Info)
	}
}

func main() {
	// Check if token is empty
	err := checkToken()
	if err != nil {
		logger.Logger.WithField("type", "token").Fatal(err)
	}

	// Create a new discordgo session
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		logger.Logger.WithField("type", "discordgo").Fatal(err)
	}
	logger.Logger.WithField("type", "discordgo").Debug("discordgo session created")

	// Add a message handler
	dg.AddHandler(messageCreate)

	// Set intents and establish connection to Discord
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = dg.Open()
	if err != nil {
		logger.Logger.WithField("type", "discordgo").Fatal(err)
	}
	logger.Logger.WithField("type", "discordgo").Info("websocket connection to Discord established")

	logger.Logger.WithField("type", "discordgo").Info("bot is now running")

	// Catch signal [SIGINT, SIGTERM] to shut down the process
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	logger.Logger.WithField("type", "process").Warn("process is interruptted")

	// Close connection to Discord
	defer func() {
		dg.Close()
		logger.Logger.WithField("type", "discordgo").Info("websocket connection to Discord closed")
		logger.Logger.WithField("type", "process").Info("diligent-parrot is stopping, bye!")
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
		logger.Logger.WithField("type", "msg").Debug("<", m.ChannelID, ">", m.Author.ID, ": ", m.Content)

		s.ChannelMessageSend(m.ChannelID, "<@!"+m.Author.ID+"> "+"Pong!")
	}
}
