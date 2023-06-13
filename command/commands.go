package command

import flags "github.com/jessevdk/go-flags"

type ExtendedCommander interface {
	flags.Commander
	Setup(UI) error
}

type Commands struct {
	Professions  ProfessionsCommand  `command:"professions" description:"List Professions"`
	Masteries    MasteriesCommand    `command:"masteries" description:"List Masteries"`
	Achievements AchievementsCommand `command:"achievements" description:"List Achievements"`
}
