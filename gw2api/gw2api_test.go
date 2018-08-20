package gw2api_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/ablease/zinn/gw2api"
)

var _ = Describe("Gw2api", func() {
	var server *ghttp.Server
	var gw2api *Gw2Api

	BeforeEach(func() {
		server = ghttp.NewServer()
		gw2api = NewAPI(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

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
				professions, err := gw2api.Professions()
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
					professions, err := gw2api.Professions()
					Expect(err).To(HaveOccurred())
					Expect(professions).To(BeNil())
				})
			})
		})
	})

	Describe("Masteries", func() {
		Context("when requesting all masteries", func() {
			var (
				statusCode       int
				masteryIDs       []int
				expectedMasterys []Mastery
				firstMastery     Mastery
				secondMastery    Mastery
			)

			BeforeEach(func() {
				statusCode = http.StatusOK
				masteryIDs = []int{}
				expectedMasterys = []Mastery{
					Mastery{Name: "foo"},
					Mastery{Name: "bar"},
				}
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/masteries"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &masteryIDs),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/masteries/1"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &firstMastery),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/masteries/2"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &secondMastery),
					),
				)
			})

			It("should return the list of masteries", func() {
				masteryIDs = []int{1, 2}
				firstMastery = Mastery{Name: "foo"}
				secondMastery = Mastery{Name: "bar"}

				response, err := gw2api.Masteries()
				Expect(err).NotTo(HaveOccurred())
				Expect(response).To(Equal(expectedMasterys))
			})
		})

		Context("when fetching masteries fails", func() {
			Context("due to a malformed response when fetching mastery ids", func() {
				var statusCode int
				var masts string
				BeforeEach(func() {
					statusCode = http.StatusOK
					masts = ""
					server.AppendHandlers(ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/masteries"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &masts),
					))
				})

				It("throws a error", func() {
					masts = "not valid json"
					masteries, err := gw2api.Masteries()
					Expect(err).To(HaveOccurred())
					Expect(masteries).To(BeNil())
				})
			})

			Context("due to a malformed response when fetching a mastery", func() {
				var statusCode int
				var masts string
				var masteryIDs []int
				BeforeEach(func() {
					statusCode = http.StatusOK
					masts = ""
					masteryIDs = []int{}
					server.AppendHandlers(ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/masteries"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &masteryIDs),
					),
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", "/v2/masteries/1"),
							ghttp.RespondWithJSONEncodedPtr(&statusCode, &masts),
						))
				})

				It("throws a error", func() {
					masteryIDs = []int{1}
					masts = "not valid json"
					masteries, err := gw2api.Masteries()
					Expect(err).To(HaveOccurred())
					Expect(masteries).To(BeNil())
				})
			})
		})
	})
})
