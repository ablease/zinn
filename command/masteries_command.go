package command

import (
	"strings"

	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . MasteriesClient
type MasteriesClient interface {
	Masteries() ([]string, error)
}

type MasteriesCommand struct {
	UI     UI
	Client MasteriesClient
}

func (m *MasteriesCommand) Setup(ui UI) error {
	m.UI = ui
	m.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}

func (m *MasteriesCommand) Execute(args []string) error {
	masts, err := m.Client.Masteries()
	if err != nil {
		return err
	}

	data := strings.Join(masts, " ")
	m.UI.DisplayText(data)
	return nil
}
