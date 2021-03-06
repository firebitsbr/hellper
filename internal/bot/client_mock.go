package bot

import (
	"context"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func NewClientMock() *ClientMock {
	return new(ClientMock)
}

func (mock *ClientMock) PostMessage(
	channel string, options ...slack.MsgOption,
) (string, string, error) {
	args := mock.Called(channel, options)
	return args.String(0), args.String(1), args.Error(2)
}

func (mock *ClientMock) CreateChannel(name string) (*slack.Channel, error) {
	var (
		args   = mock.Called(name)
		result = args.Get(0)
	)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*slack.Channel), args.Error(1)
}

func (mock *ClientMock) InviteUserToChannel(channelID string, userID string) (*slack.Channel, error) {
	var (
		args   = mock.Called(channelID, userID)
		result = args.Get(0)
	)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*slack.Channel), args.Error(1)
}

func (mock *ClientMock) ListPins(channelID string) ([]slack.Item, *slack.Paging, error) {
	var (
		args   = mock.Called(channelID)
		result = args.Get(0)
		page   = args.Get(1)
		items  []slack.Item
		paging *slack.Paging
	)
	if result != nil {
		items = result.([]slack.Item)
	}
	if page != nil {
		paging = page.(*slack.Paging)
	}
	return items, paging, args.Error(2)
}

func (mock *ClientMock) GetUserInfo(userID string) (*slack.User, error) {
	var (
		args   = mock.Called(userID)
		result = args.Get(0)
	)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*slack.User), args.Error(1)
}

func (mock *ClientMock) SetChannelTopic(channelID, topic string) (string, error) {
	args := mock.Called(channelID, topic)
	return args.String(0), args.Error(1)
}

func (mock *ClientMock) OpenDialog(string, slack.Dialog) error {
	return nil
}

func (mock *ClientMock) AddPin(string, slack.ItemRef) error {
	return nil
}

func (mock *ClientMock) ArchiveChannel(channelID string) error {
	return nil
}

func (mock *ClientMock) PostEphemeralContext(ctx context.Context, channelID string, userID string, options ...slack.MsgOption) (string, error) {
	args := mock.Called(ctx, channelID, userID, options)
	return args.String(0), args.Error(1)
}

func (mock *ClientMock) JoinConversation(string) (*slack.Channel, string, []string, error) {
	return nil, "", nil, nil
}

func (mock *ClientMock) InviteUsersToConversation(string, ...string) (*slack.Channel, error) {
	return nil, nil
}
