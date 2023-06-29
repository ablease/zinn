package command_test

import (
	"github.com/ablease/zinn/command"
	"github.com/ablease/zinn/command/commandfakes"
	"github.com/ablease/zinn/ui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Daily Reward Commands", func() {
	Describe("DailyCrafting Command", func() {
		var (
			cmd        command.DailyCraftingCommand
			testUI     *ui.UI
			err        error
			fakeClient *commandfakes.FakeApiClient
		)

		BeforeEach(func() {
			fakeClient = new(commandfakes.FakeApiClient)
			testUI = ui.NewTestUI(nil, gbytes.NewBuffer(), gbytes.NewBuffer())
			cmd = command.DailyCraftingCommand{
				BaseCommand: command.BaseCommand{
					UI:     testUI,
					Client: fakeClient,
				},
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

		Describe("fetching DailyCrafting", func() {
			Context("when the command is successfull", func() {
				BeforeEach(func() {
					fakeClient.DailyCraftingReturns([]string{"craft1", "craft2"}, nil)
				})

				It("calls the client", func() {
					_ = cmd.Execute(nil)
					Expect(fakeClient.DailyCraftingCallCount()).To(Equal(1))
				})

				It("displays the list of daily crafts", func() {
					err = cmd.Execute(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(fakeClient.DailyCraftingCallCount()).To(Equal(1))
					Expect(testUI.Out).To(gbytes.Say("craft1"))
					Expect(testUI.Out).To(gbytes.Say("craft2"))
				})
			})
		})
	})

	Describe("MapChests Command", func() {
		var (
			cmd        command.MapChestsCommand
			testUI     *ui.UI
			err        error
			fakeClient *commandfakes.FakeApiClient
		)

		BeforeEach(func() {
			fakeClient = new(commandfakes.FakeApiClient)
			testUI = ui.NewTestUI(nil, gbytes.NewBuffer(), gbytes.NewBuffer())
			cmd = command.MapChestsCommand{
				BaseCommand: command.BaseCommand{
					UI:     testUI,
					Client: fakeClient,
				},
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

		Describe("fetching mapChests", func() {
			Context("when the command is successfull", func() {
				BeforeEach(func() {
					fakeClient.MapChestsReturns([]string{"chest1", "chest2"}, nil)
				})

				It("calls the client", func() {
					_ = cmd.Execute(nil)
					Expect(fakeClient.MapChestsCallCount()).To(Equal(1))
				})

				It("displays the list of map chests", func() {
					err = cmd.Execute(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(fakeClient.MapChestsCallCount()).To(Equal(1))
					Expect(testUI.Out).To(gbytes.Say("chest1"))
					Expect(testUI.Out).To(gbytes.Say("chest2"))
				})
			})
		})
	})
})
