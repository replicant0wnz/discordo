package ui

import (
	"regexp"
	"strings"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	lua "github.com/yuin/gopher-lua"
)

var (
	boldRegex          = regexp.MustCompile(`(?ms)\*\*(.*?)\*\*`)
	italicRegex        = regexp.MustCompile(`(?ms)\*(.*?)\*`)
	underlineRegex     = regexp.MustCompile(`(?ms)__(.*?)__`)
	strikeThroughRegex = regexp.MustCompile(`(?ms)~~(.*?)~~`)
)

func parseMarkdown(md string) string {
	var res string
	res = boldRegex.ReplaceAllString(md, "[::b]$1[::-]")
	res = italicRegex.ReplaceAllString(res, "[::i]$1[::-]")
	res = underlineRegex.ReplaceAllString(res, "[::u]$1[::-]")
	res = strikeThroughRegex.ReplaceAllString(res, "[::s]$1[::-]")

	return res
}

func channelToString(c discord.Channel) string {
	var repr string

	switch c.Type {
	case discord.GuildText:
		repr = "#" + c.Name
	case discord.DirectMessage:
		rp := c.DMRecipients[0]
		repr = rp.Username + "#" + rp.Discriminator
	case discord.GroupDM:
		rps := make([]string, len(c.DMRecipients))
		for i, r := range c.DMRecipients {
			rps[i] = r.Username + "#" + r.Discriminator
		}

		repr = strings.Join(rps, ", ")
	default:
		repr = c.Name
	}

	return repr
}

func findMessageByID(ms []discord.Message, mID discord.MessageID) (int, *discord.Message) {
	for i, m := range ms {
		if m.ID == mID {
			return i, &m
		}
	}

	return -1, nil
}

func hasPermission(s *state.State, cID discord.ChannelID, p discord.Permissions) bool {
	perm, err := s.Permissions(cID, s.Ready().User.ID)
	if err != nil {
		return false
	}

	return perm&p == p
}

func returnNilLua(s *lua.LState) int {
	s.Push(lua.LNil) // Push the result
	return 1         // Number of results
}
