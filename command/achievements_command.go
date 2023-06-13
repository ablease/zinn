package command

import (
	"strconv"
	"strings"

	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . AchievementsClient
type AchievementsClient interface {
	AchievementIDs() ([]int, error)
}

type AchievementsCommand struct {
	UI     UI
	Client AchievementsClient
}

func (a *AchievementsCommand) Setup(ui UI) error {
	a.UI = ui
	a.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}

func (a *AchievementsCommand) Execute(args []string) error {
	achieves, err := a.Client.AchievementIDs()
	if err != nil {
		return err
	}

	stringData := []string{}
	for i := range achieves {
		number := achieves[i]
		text := strconv.Itoa(number)
		stringData = append(stringData, text)
	}

	data := strings.Join(stringData, " ")
	a.UI.DisplayText(data)
	return nil
}
