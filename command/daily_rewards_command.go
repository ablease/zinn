package command

import (
	"encoding/json"
	"fmt"
)

type DailyCraftingCommand struct {
	BaseCommand
}

func (d *DailyCraftingCommand) Execute(args []string) error {
	// make the call
	dailyCrafts, err := d.Client.DailyCrafting()
	if err != nil {
		return err
	}

	fmt.Printf("DailyCraftingCommand.JsonResponse: %b", d.JsonResponse)
	// decode the response
	if d.JsonResponse {
		// print the json
		fmt.Println("***** No unmarshaling")
		d.UI.DisplayText(string(dailyCrafts))
		return nil
	}

	var results []string
	err = json.Unmarshal(dailyCrafts, &results)
	if err != nil {
		return err
	}

	// print to the screen
	for _, result := range results {
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
