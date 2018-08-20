package command

import (
	"strings"

	"github.com/ablease/zinn/gw2api"
)

//go:generate counterfeiter . MasteriesClient
type MasteriesClient interface {
	Masteries() ([]gw2api.Mastery, error)
}

type MasteriesCommand struct {
	UI     UI
	Client MasteriesClient
}

func (m *MasteriesCommand) Setup(ui UI) error {
	m.UI = ui
	m.Client = gw2api.NewAPI("https://api.guildwars2.com")
	return nil
}

func (m *MasteriesCommand) Execute(args []string) error {
	masts, err := m.Client.Masteries()
	if err != nil {
		return err
	}

	// for each returned mastery create a list of their names
	masteryNames := []string{}
	for _, mastery := range masts {
		masteryNames = append(masteryNames, mastery.Name)
	}

	data := strings.Join(masteryNames, " ")
	m.UI.DisplayText(data)
	return nil
}
