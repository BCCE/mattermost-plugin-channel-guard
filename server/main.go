package main

import (
	"io/ioutil"
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	plugin.ClientMain(&guard{readFile: ioutil.ReadFile})
}
