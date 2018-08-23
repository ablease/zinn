package ui_test

import (
	. "github.com/onsi/ginkgo"
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
			ui.DisplayText("uh oh")
			Expect(out).To(Say("uh oh\n"))
		})
	})

	Describe("DisplayTable", func() {
		Context("given a matrix of strings", func() {
			var inputTable [][]string
			var firstRow []string
			var secondRow []string
			firstRow = []string{"first", "1", "a"}
			secondRow = []string{"second", "2", "b"}
			inputTable = [][]string{firstRow, secondRow}

			It("prints a table to ui.Out", func() {
				ui.DisplayTable(inputTable)
				Expect(out).To(Say("first 1 a"))
				Expect(out).To(Say("second 2 b"))
			})
		})
	})
})
