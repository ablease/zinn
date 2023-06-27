package integration

import (
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

// masteries should return a list of mastery names

var _ = Describe("masteries command", func() {
	Context("when no arguments are given", func() {
		BeforeEach(func() {
			command := exec.Command(pathToZinnCLI, "masteries")
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait(5 * time.Second)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should exit", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should print a list of mastery names", func() {
			Eventually(session).Should(gexec.Exit(0))
			Eventually(session).Should(gbytes.Say("MasteriesID  Name"))
			Eventually(session).Should(gbytes.Say("Masteries1   Exalted Lore"))
		})
	})
})
