package command_test

import (
	"github.com/ablease/zinn/command"
	"github.com/ablease/zinn/command/commandfakes"
	"github.com/ablease/zinn/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Professions Command", func() {
	var (
		cmd        command.ProfessionsCommand
		testUI     *ui.UI
		err        error
		fakeClient *commandfakes.FakeApiClient
	)

	BeforeEach(func() {
		fakeClient = new(commandfakes.FakeApiClient)
		testUI = ui.NewTestUI(nil, gbytes.NewBuffer(), gbytes.NewBuffer())
		cmd = command.ProfessionsCommand{
			UI:     testUI,
			Client: fakeClient,
		}
	})

	Describe("fetching professions", func() {
		BeforeEach(func() {
			fakeClient.ProfessionsReturns([]string{"prof1", "prof2"}, nil)
		})

		Context("when the command is successful", func() {
			It("displays the list of professions", func() {
				err = cmd.Execute(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeClient.ProfessionsCallCount()).To(Equal(1))
				Expect(testUI.Out).To(gbytes.Say("prof1 prof2"))
			})
		})
	})
})
