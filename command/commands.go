package command

import (
	"fmt"
	"github.com/ablease/zinn/api"
	flags "github.com/jessevdk/go-flags"
)

type ExtendedCommander interface {
	flags.Commander
	Setup(ui UI, json bool) error
}

type Commands struct {
	JsonResponse []bool `short:"j" long:"json" description:"Print the raw JSON response from GW2 api"`

	Professions   ProfessionsCommand   `command:"professions" description:"List Professions"`
	Masteries     MasteriesCommand     `command:"masteries" description:"List Masteries"`
	Achievements  AchievementsCommand  `command:"achievements" description:"List Achievements"`
	DailyCrafting DailyCraftingCommand `command:"daily-crafting" description:"Returns information about time-gated recipes that can be crafted in-game"`
	MapChests     MapChestsCommand     `command:"map-chests" description:"Returns information about Hero's Choice Chests that can be be acquired once a day in-game"`
	WorldBosses   WorldBossesCommand   `command:"world-bosses" description:"Returns information about scheduled World bosses in Core Tyria that reward boss chests that can be be opened once a day in-game"`
}

type BaseCommand struct {
	UI     UI
	Client ApiClient
}

func (b *BaseCommand) Setup(ui UI, jsonResponse bool) error {
	b.UI = ui
	b.JsonResponse = jsonResponse
	b.Client = api.NewClient("https://api.guildwars2.com")
	fmt.Printf("Setup, setting JsonResponse to: %b\n", jsonResponse)
	return nil
}
