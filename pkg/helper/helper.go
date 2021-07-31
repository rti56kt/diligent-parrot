package helper

import (
	"math/rand"
	"strings"

	"github.com/rti56kt/diligent-parrot/pkg/cmdlist"
	"github.com/rti56kt/diligent-parrot/pkg/cmdprefix"
	"github.com/rti56kt/diligent-parrot/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

func getThumbnail() string {
	thumbnailCandidate := []string{
		"https://i.imgur.com/CI2r3p9.gif",
		"https://i.imgur.com/PXMJvIF.gif",
	}
	randNum := rand.Float32()
	logger.Logger.WithField("type", "helper").Debug("getThumbnail rand: ", randNum)

	if randNum > 0.1 {
		return thumbnailCandidate[0]
	} else {
		return thumbnailCandidate[1]
	}
}

func getColor() int {
	colorCandidate := []int{
		0xff8d8b,
		0xfed589,
		0x88ff8a,
		0x87ffff,
		0x8bb5fe,
		0xd78cff,
		0xff8cff,
		0xff68f7,
		0xfe6cb7,
		0xff6968,
	}
	randNum := rand.Intn(10)
	logger.Logger.WithField("type", "helper").Debug("getColor rand: ", randNum)

	return colorCandidate[randNum]
}

func getCmdList() []*discordgo.MessageEmbedField {
	var embedFields []*discordgo.MessageEmbedField
	cmdAndDes := cmdlist.GetCmdAndDes()
	pre := cmdprefix.GetPrefix()

	for _, cmdObj := range cmdAndDes {
		var cmdsWithPrefix []string
		var cmdField discordgo.MessageEmbedField

		for _, cmd := range cmdObj.CommandStruct.Command {
			cmdsWithPrefix = append(cmdsWithPrefix, pre+cmd)
		}
		if cmdObj.CommandStruct.Argument == "" {
			cmdField.Name = strings.Join(cmdsWithPrefix, " | ")
		} else {
			cmdField.Name = strings.Join(cmdsWithPrefix, " | ") + " " + cmdObj.CommandStruct.Argument
		}
		cmdField.Value = cmdObj.Description
		embedFields = append(embedFields, &cmdField)
	}
	return embedFields
}

func Dealer() discordgo.MessageEmbed {
	var helpMsgEmbed discordgo.MessageEmbed

	helpMsgEmbed.Title = "Command Help"
	helpMsgEmbed.Color = getColor()

	var helpMsgEmbedThumb discordgo.MessageEmbedThumbnail
	helpMsgEmbedThumb.URL = getThumbnail()
	helpMsgEmbed.Thumbnail = &helpMsgEmbedThumb

	var helpMsgEmbedAuthor discordgo.MessageEmbedAuthor
	helpMsgEmbedAuthor.Name = "DiligentParrot"
	helpMsgEmbedAuthor.IconURL = "https://i.imgur.com/NBSGE0O.png"
	helpMsgEmbedAuthor.URL = "https://github.com/rti56kt"
	helpMsgEmbed.Author = &helpMsgEmbedAuthor

	helpMsgEmbed.Fields = getCmdList()

	var helpMsgEmbedFooter discordgo.MessageEmbedFooter
	helpMsgEmbedFooter.Text = "rti56kt"
	helpMsgEmbedFooter.IconURL = "https://avatars.githubusercontent.com/u/43367240?v=4"
	helpMsgEmbed.Footer = &helpMsgEmbedFooter

	return helpMsgEmbed
}
