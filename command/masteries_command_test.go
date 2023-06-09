package command_test

import (
	"github.com/ablease/zinn/command"
	"github.com/ablease/zinn/command/commandfakes"
	"github.com/ablease/zinn/ui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Masteries Command", func() {
	var (
		cmd        command.MasteriesCommand
		testUI     *ui.UI
		err        error
		fakeClient *commandfakes.FakeApiClient
	)

	BeforeEach(func() {
		fakeClient = new(commandfakes.FakeApiClient)
		testUI = ui.NewTestUI(nil, gbytes.NewBuffer(), gbytes.NewBuffer())
		cmd = command.MasteriesCommand{
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

	Describe("fetching masteries", func() {
		Context("when the command is successful", func() {
			BeforeEach(func() {
				fakeClient.MasteriesReturns([]string{"a mastery", "another mastery"}, nil)
			})

			It("calls the client", func() {
				_ = cmd.Execute(nil)
				Expect(fakeClient.MasteriesCallCount()).To(Equal(1))
			})

			It("displays the list of masteries", func() {
				err = cmd.Execute(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeClient.MasteriesCallCount()).To(Equal(1))
				Expect(testUI.Out).To(gbytes.Say("a mastery another mastery"))
			})
		})
	})
})
