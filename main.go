package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ablease/zinn/command"
	"github.com/ablease/zinn/ui"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	var commands command.Commands
	parser := flags.NewParser(&commands, flags.HelpFlag)
	parser.CommandHandler = handleCommand
	_, err := parser.ParseArgs(os.Args[1:])

	if len(os.Args[1:]) < 1 || os.Args[1] == "help" {
		parser.WriteHelp(os.Stdout)
	}

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
}

func handleCommand(cmd flags.Commander, args []string) error {
	commandUI := ui.NewUI()

	if extendedCmd, ok := cmd.(command.ExtendedCommander); ok {
		err := extendedCmd.Setup(commandUI)
		if err != nil {
			return handleError(err, commandUI)
		}
		return handleError(extendedCmd.Execute(args), commandUI)
	}

	return nil
}

func handleError(passedErr error, commandUI *ui.UI) error {
	if passedErr == nil {
		return nil
	}

	commandUI.DisplayError(passedErr)

	return errors.New("command failed")
}
