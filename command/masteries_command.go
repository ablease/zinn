package command

import (
	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . MasteriesClient
type MasteriesClient interface {
	Masteries(ids []int) ([]api.Mastery, error)
	GetMasteryIDs() ([]int, error)
}

type MasteriesCommand struct {
	UI         UI
	Client     MasteriesClient
	MasteryIDs []int `short:"i" long:"id" description:"Get a specific Mastery by ID"`
}

func (m *MasteriesCommand) Setup(ui UI) error {
	m.UI = ui
	m.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}

func (m *MasteriesCommand) Execute(args []string) error {
	var ids []int
	var err error
	if len(m.MasteryIDs) == 0 {
		ids, err = m.Client.GetMasteryIDs()
		if err != nil {
			return err
		}
	} else {
		ids = m.MasteryIDs
	}

	masts, err := m.Client.Masteries(ids)
	if err != nil {
		return err
	}

	table := [][]string{}
	headerRow := []string{"ID", "Name", "Requirement", "Order", "Region"}
	table = append(table, headerRow)

	for _, mastery := range masts {
		row := []string{}
		row = append(
			row,
			intToString(mastery.ID),
			mastery.Name,
			mastery.Requirement,
			intToString(mastery.Order),
			mastery.Region,
		)
		table = append(table, row)
	}

	m.UI.DisplayNonWrappingTable("", table, 2)

	m.UI.DisplayText("\nFor more information about specific masteries see 'zinn mastery'")
	return nil
}

func (m *MasteriesCommand) Usage() string {
	return "--id=ID\n\n ID is a mastery ID.\n If ID is omitted all masteries are returned."
}

func (m *MasteriesCommand) Examples() string {
	return `
zinn masteries			     # returns all masteries
zinn masteries -i=1	     # returns the Exalted Lore mastery
zinn masteries --id=1 --id=2 # returns the Exhalted Lore and Itzel Lore Mastery
`
}
