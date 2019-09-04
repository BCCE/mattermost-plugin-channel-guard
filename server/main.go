package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&guard{})
}
