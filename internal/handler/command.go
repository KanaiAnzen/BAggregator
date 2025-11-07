package handler

import "fmt"

func (c *Commands) Run(s *State, cmd Command) error {
	if f, ok := c.cmds[cmd.name]; ok {
		return f(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.cmds[name] = f
}
