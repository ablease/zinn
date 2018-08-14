package command

import (
	"strconv"
	"strings"

	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . MasteriesClient
type MasteriesClient interface {
	Masteries() ([]int, error)
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

	ms := []string{}
	for _, mastery := range masts {
		i := strconv.Itoa(mastery)
		ms = append(ms, i)
	}

	data := strings.Join(ms, " ")
	m.UI.DisplayText(data)
	return nil
}
