# Channel Guard Plugin

Use this plugin to make channels read-only to some and writeable to other users. Channel Admins, Team admin, bots and system are all allowed to post. 

## Configuration

1. Go to **System Console > Plugins > Management** and click **Enable** to enable the Channel Guard plugin.

2. Modify your `config.json` file to include your Guards, under the `PluginSettings`. See below for an example of what this should look like.

## Usage

To configure the Guard, edit your `config.json` file with the following format:

```
"plugins": {
	"com.mattermost.channel-guard": {
		"guards":[
			{
			"TeamName": "your-team-name",
			"ChannelName": "channel_1",
			"Allowed" : ["user_1", "user_1"]
			}
		]
	}
}
```

where

- **TeamName**: The team for which the guard will match against. Must be the team handle used in the URL, in lowercase. For example, in the following URL the **TeamName** value is `my-team`: https://example.com/my-team/channels/my-channel
- **ChannelName**:  The channel that is under guard. Must be the channel handle used in the URL, in lowercase. For example, in the following URL the **channel name** value is `my-channel`: https://example.com/my-team/channels/my-channel
- **Allowed**: List of Mattermost Usersnames that can post in this channel.

## Development
For additional information on developing plugins, refer to [our plugin developer documentation](https://developers.mattermost.com/extend/plugins/).
