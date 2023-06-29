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

type WorldBossesCommand struct {
	BaseCommand
}

func (w *WorldBossesCommand) Execute(args []string) error {
	bosses, err := w.Client.WorldBosses()
	if err != nil {
		return err
	}

	for _, result := range bosses {
		w.UI.DisplayText(result)
	}
	return nil
}
