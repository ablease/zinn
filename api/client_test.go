package api_test

import (
	"net/http"

	. "github.com/ablease/zinn/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {
	var server *ghttp.Server
	var client *Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})
	// Error scenarios.. when NewRequest returns an error.. (url is malformed?)
	// when client can't make the request (httpClient.Do)
	// when json unmarshal fails
	Describe("Professions", func() {
		Context("when requesting all professions", func() {
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

			It("should return the returned professions", func() {
				profs = []string{"prof1", "prof2"}
				professions, err := client.Professions()
				Expect(err).NotTo(HaveOccurred())
				Expect(professions).To(Equal(profs))
			})
		})

		Context("when fetching professions fails", func() {
			Context("due to a malformed response", func() {
				var statusCode int
				var profs string
				BeforeEach(func() {
					statusCode = http.StatusOK
					profs = ""
					server.AppendHandlers(ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/professions"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &profs),
					))
				})

				It("throws a error", func() {
					profs = "not valid json"
					professions, err := client.Professions()
					Expect(err).To(HaveOccurred())
					Expect(professions).To(BeNil())
				})
			})
		})
	})
})
