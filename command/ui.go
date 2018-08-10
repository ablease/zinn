package command

type UI interface {
	DisplayText(data string)
	DisplayError(err error)
}
