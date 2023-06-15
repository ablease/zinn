package ui_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "github.com/ablease/zinn/ui"
)

var _ = Describe("UI", func() {
	var (
		ui      *UI
		out     *Buffer
		errBuff *Buffer
	)

	BeforeEach(func() {
		ui = NewUI()

		out = NewBuffer()
		ui.Out = out
		errBuff = NewBuffer()
		ui.Err = errBuff
	})

	Describe("DisplayText", func() {
		It("displays the given string to ui.Out with a newline", func() {
			ui.DisplayText("boop")
			Expect(out).To(Say("boop\n"))
		})
	})

	Describe("DisplayError", func() {
		It("displays the given error to ui.Err with a newline", func() {
			ui.DisplayError(errors.New("uh oh"))
			Expect(errBuff).To(Say("uh oh\n"))
		})
	})

	Describe("DisplayNonWrappingTable", func() {
		It("displays the provided table with prefix", func() {
			testdata := [][]string{
				{"row one", "column2", "column3"},
				{"row two", "column2", "column3"},
				{"row three", "column2", "column3"},
			}
			ui.DisplayNonWrappingTable("prefix", testdata, 2)
			Eventually(out, 3).Should(Say("prefixrow one    column2  column3"))
			Eventually(out, 3).Should(Say("prefixrow two    column2  column3"))
			Eventually(out, 3).Should(Say("prefixrow three  column2  column3"))
		})
	})
})
