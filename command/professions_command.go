package command

import (
	"strings"

	"github.com/ablease/zinn/api"
)

//go:generate counterfeiter . ProfessionClient
type ProfessionClient interface {
	Professions() ([]string, error)
}

type ProfessionsCommand struct {
	UI     UI
	Client ProfessionClient
}

func (p *ProfessionsCommand) Setup(ui UI) error {
	p.UI = ui
	p.Client = api.NewZinnClient("https://api.guildwars2.com")
	return nil
}

func (p *ProfessionsCommand) Execute(args []string) error {
	profs, err := p.Client.Professions()
	if err != nil {
		return err
	}

	data := strings.Join(profs, " ")
	p.UI.DisplayText(data)
	return nil
}
