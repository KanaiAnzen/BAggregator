package handler

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("login requires one argument")
	}

	username := cmd.arg[0]
	if s.cfgp == nil {
		return fmt.Errorf("config is not initialized")
	}

	err := s.cfgp.SetUser(username)
	if err != nil {
		return fmt.Errorf("setting username failure")
	}

	fmt.Printf("User has been set to %q\n", username)

	return nil
}
