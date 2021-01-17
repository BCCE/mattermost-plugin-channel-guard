package main

import (
	"io/ioutil"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func main() {
	plugin.ClientMain(&guard{readFile: ioutil.ReadFile})
}
