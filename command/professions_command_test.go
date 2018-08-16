package command_test

import (
	"errors"

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

	Describe("Setup", func() {
		It("Sets the UI object", func() {
			err = cmd.Setup(testUI)
			Expect(err).ToNot(HaveOccurred())
			Expect(cmd.UI).To(Equal(testUI))
		})

		It("Sets the Client object", func() {
			err = cmd.Setup(testUI)
			Expect(err).ToNot(HaveOccurred())
			Expect(cmd.Client).ToNot(Equal(nil))
		})
	})

	Describe("fetching professions", func() {
		Context("when the command is successful", func() {
			BeforeEach(func() {
				fakeClient.ProfessionsReturns([]string{"prof1", "prof2"}, nil)
			})

			It("displays the list of professions", func() {
				err = cmd.Execute(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeClient.ProfessionsCallCount()).To(Equal(1))
				Expect(testUI.Out).To(gbytes.Say("prof1 prof2"))
			})
		})

		Context("when the command is unsuccessful", func() {
			Context("due to a client error", func() {
				BeforeEach(func() {
					fakeClient.ProfessionsReturns(nil, errors.New("oops"))
				})

				It("displays the error to the ui", func() {
					err = cmd.Execute(nil)
					Expect(err).To(HaveOccurred())
					Expect(fakeClient.ProfessionsCallCount()).To(Equal(1))
					Expect(err.Error()).To(Equal("oops"))
				})
			})
		})
	})
})
