package integration

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

// professions should return a list of professions

var _ = Describe("professions command", func() {
	Context("when no arguments are given", func() {
		BeforeEach(func() {
			command := exec.Command(pathToZinnCLI, "professions")
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should exit", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should print a list of professions", func() {
			Eventually(session).Should(gexec.Exit(0))
			Eventually(session).Should(gbytes.Say("Guardian"))
			Eventually(session).Should(gbytes.Say("Warrior"))
			Eventually(session).Should(gbytes.Say("Engineer"))
			Eventually(session).Should(gbytes.Say("Ranger"))
			Eventually(session).Should(gbytes.Say("Thief"))
			Eventually(session).Should(gbytes.Say("Elementalist"))
			Eventually(session).Should(gbytes.Say("Mesmer"))
			Eventually(session).Should(gbytes.Say("Necromancer"))
			Eventually(session).Should(gbytes.Say("Revenant"))
		})
	})
})
