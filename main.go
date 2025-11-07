package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TadaTeki/BAggregator/internal/config"
	"github.com/TadaTeki/BAggregator/internal/handler"
)

func main() {
	cfgp, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// init
	sta := handler.NewState(cfgp)

	cmds := handler.NewCommands()
	cmds.Register("login", handler.HandlerLogin)

	//
	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("Argument necessary")
	}

	fmt.Println(len(arguments))
	for i := 0; i < len(arguments); i++ {
		fmt.Println(i, arguments[i])
	}

	arg := []string{}
	if len(arguments) > 2 {
		arg = arguments[2:]
	}
	cmd := handler.NewCommand(arguments[1], arg)

	err2 := cmds.Run(sta, *cmd)
	if err2 != nil {
		log.Fatalf("error executing command: %v", err2)
	}

}
