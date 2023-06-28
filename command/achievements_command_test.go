package command_test

import (
	"github.com/ablease/zinn/api"
	"github.com/ablease/zinn/command"
	"github.com/ablease/zinn/command/commandfakes"
	"github.com/ablease/zinn/ui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Achievements Command", func() {
	var (
		cmd        command.AchievementsCommand
		testUI     *ui.UI
		err        error
		fakeClient *commandfakes.FakeApiClient
	)

	BeforeEach(func() {
		fakeClient = new(commandfakes.FakeApiClient)
		testUI = ui.NewTestUI(nil, gbytes.NewBuffer(), gbytes.NewBuffer())
		cmd = command.AchievementsCommand{
			UI:     testUI,
			Client: fakeClient,
		}
	})

	Describe("Setup", func() {
		It("Sets up the UI object", func() {
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

	Describe("fetching Achievements", func() {
		Context("when the command is successfull", func() {
			BeforeEach(func() {
				fakeClient.AchievementIDsReturns([]int{1, 2}, nil)
			})

			It("calls the client", func() {
				_ = cmd.Execute(nil)
				Expect(fakeClient.AchievementIDsCallCount()).To(Equal(1))
			})

			It("displays the list of achievements", func() {
				err = cmd.Execute(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeClient.AchievementIDsCallCount()).To(Equal(1))
				Expect(testUI.Out).To(gbytes.Say("1 2"))
			})
		})

		Context("when one ID is provided", func() {
			BeforeEach(func() {
				fakeClient.AchievementsReturns([]api.Achievement{{Name: "achievement1"}}, nil)
			})

			It("returns the request Achievement", func() {
				cmd.AchievementIDs = []int{1}
				_ = cmd.Execute(nil)
				Expect(fakeClient.AchievementIDsCallCount()).To(Equal(0))
				Expect(fakeClient.AchievementsCallCount()).To(Equal(1))
				Expect(fakeClient.AchievementsArgsForCall(0)).To(Equal([]int{1}))
				Expect(testUI.Out).To(gbytes.Say("achievement1"))
			})
		})
	})
})
