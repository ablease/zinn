package command

type DailyCraftingCommand struct {
	BaseCommand
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

type MapChestsCommand struct {
	BaseCommand
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
