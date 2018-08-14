package integration

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

// masteries should return a list of masteries
// masteries should accept an argument which is the mastery (name/id)
// and returnt the requested mastery

var _ = Describe("masteries command", func() {
	Context("when no arguments are given", func() {
		BeforeEach(func() {
			command := exec.Command(pathToZinnCLI, "masteries")
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should exit", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should print a list of mastery ids", func() {
			Eventually(session).Should(gexec.Exit(0))
			Eventually(session).Should(gbytes.Say("1 2 3 4 5 6 8 12 13 14 15 16 17 18 19 20"))
		})
	})
})
