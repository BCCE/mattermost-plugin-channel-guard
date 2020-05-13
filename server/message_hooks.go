package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// MessageWillBePosted intercepts each post to either reject or allow based in function criteria
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	if len(p.config.ReadOnlyChannels) != 0 {
		if p.readonlyChannelChecker(post) {
			if post.IsSystemMessage() {
				return post, ""
			}

			if p.channelRoleChecker(post) {
				return post, ""
			}

			postuser, _ := p.API.GetUser(post.UserId)

			if postuser.IsBot == true {
				return post, ""
			}

			if len(p.config.AllowedUsers) != 0 {
				if p.allowedUsersChecker(post) {
					return post, ""
				}
			}
			p.API.SendEphemeralPost(post.UserId, &model.Post{
				UserId:    p.botUserID,
				ChannelId: post.ChannelId,
				Message:   p.config.ChannelGuardMsg,
			})

			// username & channelname are not needed, only used for logging on line 44
			username, _ := p.API.GetUser(post.UserId)
			channelname, _ := p.API.GetChannel(post.ChannelId)

			msg := fmt.Sprintf("%s (UserID: %s) attempted to post in the Read-Only Channel %s (ChannelID: %s)", username.Username, post.UserId, channelname.Name, post.ChannelId)
			return nil, msg
		}
	}

	return post, ""
}

// roleChecker checks role of user to determine if they are a channel admin
func (p *Plugin) channelRoleChecker(post *model.Post) bool {
	channelMember, _ := p.API.GetChannelMember(post.ChannelId, post.UserId)
	channelRoles := channelMember.GetRoles()

	for _, role := range channelRoles {
		if role == "channel_admin" {
			return true
		}
	}
	return false
}

// readonlyChannelChecker checks read-only channels from config and compares the post channel
func (p *Plugin) readonlyChannelChecker(post *model.Post) bool {
	readonlyChannels := strings.Split(p.config.ReadOnlyChannels, ", ")
	team, _ := p.API.GetTeamByName(p.config.TeamName)

	for _, channel := range readonlyChannels {
		rochan, _ := p.API.GetChannelByName(team.Id, channel, false)
		if post.ChannelId == rochan.Id {
			return true
		}
	}
	return false
}

// allowedUsersChecker checks allowed users from config and compares the post user
func (p *Plugin) allowedUsersChecker(post *model.Post) bool {
	allowedusers := strings.Split(p.config.AllowedUsers, ", ")

	for _, user := range allowedusers {
		alloweduser, _ := p.API.GetUserByUsername(user)
		if post.UserId == alloweduser.Id {
			return true
		}
	}
	return false
}
