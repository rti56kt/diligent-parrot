package main

import (
	"errors"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rti56kt/diligent-parrot/pkg/cmdlist"
	"github.com/rti56kt/diligent-parrot/pkg/cmdprefix"
	"github.com/rti56kt/diligent-parrot/pkg/helper"
	"github.com/rti56kt/diligent-parrot/pkg/i18n"
	"github.com/rti56kt/diligent-parrot/pkg/ifconfigme"
	"github.com/rti56kt/diligent-parrot/pkg/logger"
	"github.com/rti56kt/diligent-parrot/pkg/msgparser"
	"github.com/rti56kt/diligent-parrot/pkg/msgresponder"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken   string = os.Getenv("DISCORD_BOT_TOKEN")
	logVerbose bool
)

func init() {
	flag.BoolVar(&logVerbose, "v", false, "Log verbose (false:INFO true:DEBUG)")
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

	// Detect prefix to determine if the msg is a cmd
	prefix := cmdprefix.GetPrefix()
	if strings.HasPrefix(m.Content, prefix) {
		msgparser.CmdPreprocess(m, prefix)
	} else {
		if resp, exist := msgresponder.GetKeywordResp(m.Content); exist {
			resp = msgparser.GetAuthorTag(m) + " " + resp
			s.ChannelMessageSend(m.ChannelID, resp)
		} else {
			return
		}
	}

	dbgMsg := msgparser.GetCmdDetail(m)
	logger.Logger.WithField("type", "msg").Debug(dbgMsg)

	// Deal with cmd
	cmdAndArgs := msgparser.GetCmdAndArgs(m)
	if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Ping) {
		// "ping" cmd
		resp := msgparser.GetAuthorTag(m) + " " + i18n.AllLocale[i18n.GetCurrentLocale()].PING.PONG
		s.ChannelMessageSend(m.ChannelID, resp)
	} else if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Ifconfigme) {
		// "ifconfigme" cmd
		resp := ifconfigme.Dealer()
		s.ChannelMessageSend(m.ChannelID, resp)
	} else if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Help) {
		// "help" cmd
		resp := helper.Dealer()
		s.ChannelMessageSendEmbed(m.ChannelID, &resp)
	} else if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Set) {
		// "set" cmd
		resp := msgresponder.Dealer(m)
		s.ChannelMessageSend(m.ChannelID, resp)
	} else if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Locale) {
		// "locale" cmd
		resp := i18n.Dealer(m)
		s.ChannelMessageSendComplex(m.ChannelID, &resp)
	} else if cmdlist.CheckCmdByFullName(cmdAndArgs[0], cmdlist.Prefix) {
		// "prefix" cmd
		resp := cmdprefix.Dealer(s, m)
		s.ChannelMessageSend(m.ChannelID, resp)
	}
}
