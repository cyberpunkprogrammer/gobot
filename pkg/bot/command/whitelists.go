package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/whitelist"
)

type whitelists struct {
	command
}

func init() {
	whitelists := whitelists{command{
		name:        "whitelists",
		description: "Lists bot whitelists.",
		permissions: []int{discordgo.PermissionManageServer},
	}}
	executables = append(executables, &whitelists)
}

func (a *whitelists) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	channel := message.ChannelID

	if len(whitelist.Whitelists) == 0 {
		session.ChannelMessageSend(channel, "<@"+message.Author.ID+"> no whitelists have been set.")
		return
	}

	output := "**Bot Whitelists**\n"
	output += ">>> "
	for _, whitelist := range whitelist.Whitelists {
		output += "*" + whitelist.Text + "*\n"
	}

	session.ChannelMessageSend(channel, output)
}
