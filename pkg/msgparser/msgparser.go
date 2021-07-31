package msgparser

import (
	"strings"

	"github.com/rti56kt/diligent-parrot/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

func GetCmdAndArgs(m *discordgo.MessageCreate) []string {
	cmdAndArgs := strings.Split(m.Content, " ")
	return cmdAndArgs
}

func GetAuthorTag(m *discordgo.MessageCreate) string {
	authorTag := "<@!" + m.Author.ID + ">"
	return authorTag
}

func GetCmdDetail(m *discordgo.MessageCreate) string {
	cmdDetail := "<" + m.ChannelID + ">" + m.Author.ID + ": " + m.Content
	return cmdDetail
}

func CmdPreprocess(m *discordgo.MessageCreate, prefix string) {
	m.Content = strings.ToLower(m.Content)
	m.Content = strings.TrimPrefix(m.Content, prefix)
}

func IsAuthorAdmin(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	for _, roleID := range m.Member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			logger.Logger.WithField("type", "discordgo").Error(err)
		}
		logger.Logger.WithField("type", "parser").Debug(role.Permissions)
		logger.Logger.WithField("type", "parser").Debug((role.Permissions >> 3) & 1)
		if (role.Permissions>>3)&1 == 1 {
			return true
		}
	}
	return false
}

func IsAuthorOwner(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		logger.Logger.WithField("type", "discordgo").Error(err)
	}
	if m.Author.ID == guild.OwnerID {
		return true
	}
	return false
}
