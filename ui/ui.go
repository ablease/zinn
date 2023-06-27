package ui

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/lunixbochs/vtclean"
	runewidth "github.com/mattn/go-runewidth"
)

type UI struct {
	// Out is the output buffer that prints to stdout
	Out io.Writer

	// Err is the output buffer that prints to stderr
	Err io.Writer

	terminalLock *sync.Mutex
}

// NewUI will return a UI object where Out is set to STDOUT
// and Err is set to STDERR
func NewUI() *UI {
	return &UI{
		Out:          os.Stdout,
		Err:          os.Stderr,
		terminalLock: &sync.Mutex{},
	}
}

// NewTestUI will return a UI object where Out and Err are customizable.
func NewTestUI(in io.Reader, out io.Writer, err io.Writer) *UI {
	return &UI{
		Out:          out,
		Err:          err,
		terminalLock: &sync.Mutex{},
	}
}

// DisplayText outputs the given string to ui.Out.
func (ui *UI) DisplayText(data string) {
	ui.terminalLock.Lock()
	defer ui.terminalLock.Unlock()

	io.WriteString(ui.Out, data)
	fmt.Fprintf(ui.Out, "\n")
}

// DisplayError outputs the given error to ui.Err.
func (ui *UI) DisplayError(message error) {
	ui.terminalLock.Lock()
	defer ui.terminalLock.Unlock()

	_, err := io.WriteString(ui.Err, message.Error())
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(ui.Err, "\n")
	if err != nil {
		panic(err)
	}
}

// DisplayJSON

// DisplayNonWrappingTable draws a table from a 2d array of strings to ui.Out.
// Prefix will be prepended to each row. Padding adds spaces between columns
func (ui *UI) DisplayNonWrappingTable(prefix string, table [][]string, padding int) {
	ui.terminalLock.Lock()
	defer ui.terminalLock.Unlock()

	if len(table) == 0 {
		return
	}

	var columnPadding []int

	rows := len(table)
	columns := len(table[0])
	for col := 0; col < columns; col++ {
		var max int
		for row := 0; row < rows; row++ {
			if strLen := wordSize(table[row][col]); max < strLen {
				max = strLen
			}
		}
		columnPadding = append(columnPadding, max+padding)
	}

	for row := 0; row < rows; row++ {
		fmt.Fprint(ui.Out, prefix)
		for col := 0; col < columns; col++ {
			data := table[row][col]
			var addedPadding int
			if col+1 != columns {
				addedPadding = columnPadding[col] - wordSize(data)
			}
			fmt.Fprintf(ui.Out, "%s%s", data, strings.Repeat(" ", addedPadding))
		}
		fmt.Fprintf(ui.Out, "\n")
	}
}

func wordSize(str string) int {
	cleanStr := vtclean.Clean(str, false)
	return runewidth.StringWidth(cleanStr)
}
