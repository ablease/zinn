package command

import flags "github.com/jessevdk/go-flags"

type ExtendedCommander interface {
	flags.Commander
	Setup(UI) error
}

type Commands struct {
	Professions ProfessionsCommand `command:"professions" description:"List Professions"`
}
