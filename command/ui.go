package command

type UI interface {
	DisplayText(data string)
	DisplayError(err error)
	DisplayNonWrappingTable(prefix string, table [][]string, padding int)
}
