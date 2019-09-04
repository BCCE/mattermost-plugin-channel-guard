package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

const message = "This channel is under guard. You do not have the permissions to post. Please contact the system administrators if you beleive this is incorrect"

func (p *guard) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {

	Teams, appErr := p.API.GetTeams()
	if appErr != nil {
		return nil, "Failed to get teams"
	}

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

	for _, guard := range p.getGuards() {

		users, _ := p.API.GetUsersByUsernames(guard.Allowed)

		for _, team := range Teams {

			if team.Name == guard.TeamName {

				if p.teamRoleChecker(post, team) {
					return post, ""
				}

				returnedPost := p.checker(post, team, users, guard)
				if returnedPost != nil {
					return post, ""
				}

				str := fmt.Sprintf("%s attempted to post in channel %s", post.UserId, post.ChannelId)

				return nil, str

			}
		}
	}

	return post, ""
}

func (p *guard) teamRoleChecker(post *model.Post, team *model.Team) bool {

	teamMember, _ := p.API.GetTeamMember(team.Id, post.UserId)

	teamRoles := teamMember.GetRoles()

	for _, a := range teamRoles {
		if a == "team_admin" {
			return true
		}
	}

	return false

}

func (p *guard) channelRoleChecker(post *model.Post) bool {

	channelMember, _ := p.API.GetChannelMember(post.ChannelId, post.UserId)

	channelRoles := channelMember.GetRoles()

	for _, a := range channelRoles {
		if a == "channel_admin" {
			return true
		}
	}

	return false

}

func (p *guard) checker(post *model.Post, team *model.Team, users []*model.User, guard *ConfigGuard) *model.Post {

	guardChannel, _ := p.API.GetChannelByName(team.Id, guard.ChannelName, false)

	if post.ChannelId == guardChannel.Id {
		if len(users) == 0 {
			p.API.SendEphemeralPost(post.UserId, &model.Post{
				UserId:    p.botUserID,
				ChannelId: guardChannel.Id,
				Message:   message,
			})

			return nil
		}

		for _, user := range users {

			if post.UserId == user.Id {
				return post
			}

			p.API.SendEphemeralPost(post.UserId, &model.Post{
				UserId:    p.botUserID,
				ChannelId: guardChannel.Id,
				Message:   message,
			})

			return nil

		}
	}

	return post
}
