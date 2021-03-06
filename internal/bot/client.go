package bot

import (
	"context"

	"github.com/slack-go/slack"
)

type Client interface {
	PostEphemeralContext(context.Context, string, string, ...slack.MsgOption) (string, error)
	PostMessage(string, ...slack.MsgOption) (string, string, error)
	CreateChannel(string) (*slack.Channel, error)
	InviteUserToChannel(string, string) (*slack.Channel, error)
	ListPins(string) ([]slack.Item, *slack.Paging, error)
	GetUserInfo(string) (*slack.User, error)
	SetChannelTopic(string, string) (string, error)
	OpenDialog(string, slack.Dialog) error
	AddPin(string, slack.ItemRef) error
	ArchiveChannel(channelID string) error
	JoinConversation(string) (*slack.Channel, string, []string, error)
	InviteUsersToConversation(string, ...string) (*slack.Channel, error)
}
