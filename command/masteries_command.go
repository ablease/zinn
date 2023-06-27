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
	UI     UI
	Client MasteriesClient
}

func (m *MasteriesCommand) Setup(ui UI) error {
	m.UI = ui
	m.Client = api.NewClient("https://api.guildwars2.com")
	return nil
}

func (m *MasteriesCommand) Execute(args []string) error {
	ids, err := m.Client.GetMasteryIDs()
	if err != nil {
		return err
	}

	masts, err := m.Client.Masteries(ids)
	if err != nil {
		return err
	}

	headerRow := []string{"ID", "Name", "Requirement", "Order", "Region"}
	// for the number of masterys, create that many rows for the number of mastery fields create that many cols
	// numRows := len(masts)
	// numCols := reflect.ValueOf(api.Mastery{}).NumField()
	table := [][]string{}
	table = append(table, headerRow)

	// we have the table, lets set the column headers (first row is field names)
	// get the field names for the struct
	//fields := reflect.VisibleFields(reflect.TypeOf(struct{ api.Mastery }{}))
	//headers := []string{}
	//for _, field := range fields {
	//	headers = append(headers, field.Name)
	//}
	//table = append(table, headers)
	// [[Mastery ID Name Requirement Order Background Region Levels] [      ]]

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
	return nil
}
