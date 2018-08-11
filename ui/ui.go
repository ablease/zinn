package ui

import (
	"fmt"
	"io"
	"os"
	"sync"
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
		// In:           os.Stdin,
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

	_, err := io.WriteString(ui.Out, data)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(ui.Out, "\n")
	if err != nil {
		panic(err)
	}
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
