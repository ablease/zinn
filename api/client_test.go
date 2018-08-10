package api_test

import (
	"net/http"

	. "github.com/ablease/zinn/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("A Zinn Client", func() {
	var server *ghttp.Server
	var client *ZinnClient

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewZinnClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("fetching professions", func() {
		var statusCode int
		var profs []string
		BeforeEach(func() {
			statusCode = http.StatusOK
			profs = []string{}
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/professions"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &profs),
			))
		})

		Context("when requesting all professions", func() {
			Context("when the response is succesful", func() {
				BeforeEach(func() {
					profs = []string{"prof1", "prof2"}
				})

				It("should return the returned professions", func() {
					professions, err := client.Professions()
					Expect(err).NotTo(HaveOccurred())
					Expect(professions).Should(Equal(profs))
				})
			})
		})
	})
})
