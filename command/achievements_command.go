package command

import (
	"strconv"
	"strings"

	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . AchievementsClient
type AchievementsClient interface {
	AchievementIDs() ([]int, error)
	Achievements(ids []int) ([]api.Achievement, error)
}

type AchievementsCommand struct {
	UI             UI
	Client         AchievementsClient
	AchievementIDs []int `long:"id" description:"Get a specific Achievement by ID"`
}

// Setup sets the UI object, and the URL for the api Client
func (a *AchievementsCommand) Setup(ui UI) error {
	a.UI = ui
	a.Client = api.NewClient("https://api.guildwars2.com/v2/achievements")
	return nil
}

func (a *AchievementsCommand) Execute(args []string) error {
	if len(a.AchievementIDs) == 0 {
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
	} else {
		achieves, err := a.Client.Achievements(a.AchievementIDs)
		if err != nil {
			return err
		}
		table := [][]string{}
		headerRow := []string{"ID", "Name", "Description", "Requirement", "LockedText", "Type"}
		table = append(table, headerRow)

		for _, achievement := range achieves {
			row := []string{}
			row = append(row,
				intToString(achievement.ID),
				achievement.Name,
				achievement.Description,
				achievement.Requirement,
				achievement.LockedText,
				achievement.Type,
			)
			table = append(table, row)
		}
		a.UI.DisplayNonWrappingTable("", table, 2)
	}
	a.UI.DisplayText("\nFor more information about specific achievements see 'zinn mastery'")
	return nil
}
