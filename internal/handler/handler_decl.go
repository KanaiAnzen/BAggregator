package handler

import (
	"github.com/TadaTeki/BAggregator/internal/config"
)

type State struct {
	cfgp *config.Config
}

type Command struct {
	name string
	arg  []string
}

type Commands struct {
	cmds map[string]func(*State, Command) error
}

func NewState(cfg *config.Config) *State {
	return &State{cfgp: cfg}
}

func NewCommand(str string, arg []string) *Command {
	return &Command{name: str, arg: arg}
}

func NewCommands() *Commands {
	return &Commands{cmds: make(map[string]func(*State, Command) error)}
}
