package command

import (
	"strings"
)

type ProfessionsCommand struct {
	BaseCommand
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
