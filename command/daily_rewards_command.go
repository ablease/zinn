package command

import "github.com/ablease/zinn/api"

//go:generate counterfeiter . DailyCraftingClient
type DailyCraftingClient interface {
	DailyCrafting() ([]string, error)
}

type DailyCraftingCommand struct {
	UI     UI
	Client DailyCraftingClient
}

// Setup sets the UI object, and the URL for the api Client
func (d *DailyCraftingCommand) Setup(ui UI) error {
	d.UI = ui
	d.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}

func (d *DailyCraftingCommand) Execute(args []string) error {
	dailyCrafts, err := d.Client.DailyCrafting()
	if err != nil {
		return err
	}

	for _, result := range dailyCrafts {
		d.UI.DisplayText(result)
	}
	return nil
}

//go:generate counterfeiter . MapChestsClient
type MapChestsClient interface {
	MapChests() ([]string, error)
}

type MapChestsCommand struct {
	UI     UI
	Client MapChestsClient
}

func (m *MapChestsCommand) Setup(ui UI) error {
	m.UI = ui
	m.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}
func (m *MapChestsCommand) Execute(args []string) error {
	mapChests, err := m.Client.MapChests()
	if err != nil {
		return err
	}

	for _, result := range mapChests {
		m.UI.DisplayText(result)
	}
	return nil
}
