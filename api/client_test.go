package api_test

import (
	"net/http"

	. "github.com/ablease/zinn/api"
	. "github.com/onsi/ginkgo/v2"
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

	Describe("Masteries", func() {
		Context("when requesting all masteries", func() {
			var statusCode int
			var returnedMasteries []Mastery

			BeforeEach(func() {
				statusCode = http.StatusOK
				returnedMasteries = []Mastery{{
					ID:   1,
					Name: "Mastery1",
				},
					{
						ID:          2,
						Name:        "Mastery2",
						Requirement: "requirement",
					},
				}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v2/masteries", "ids=1,2"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedMasteries),
				))
			})

			It("should return the returned masterys", func() {
				masteries, err := client.Masteries([]int{1, 2})
				Expect(err).NotTo(HaveOccurred())
				Expect(masteries).To(Equal(returnedMasteries))
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
					masteries, err := client.Masteries([]int{1, 2})
					Expect(err).To(HaveOccurred())
					Expect(masteries).To(BeNil())
				})
			})
		})
	})

	Describe("Achievements", func() {
		Context("when requesting all achievements", func() {
			var (
				statusCode int
				achieveIDs []int
			)

			BeforeEach(func() {
				statusCode = http.StatusOK
				achieveIDs = []int{}
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &achieveIDs),
					),
				)
			})

			It("should return the list of achievements id's", func() {
				achieveIDs = []int{1, 2}

				response, err := client.AchievementIDs()
				Expect(err).NotTo(HaveOccurred())
				Expect(response).To(Equal([]int{1, 2}))
			})
		})

		Context("when fetching achievementIDs fails", func() {
			Context("due to a malformed response when fetching achievement ids", func() {
				var statusCode int
				var achieveIDs string
				BeforeEach(func() {
					statusCode = http.StatusOK
					achieveIDs = ""
					server.AppendHandlers(ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, &achieveIDs),
					))
				})

				It("throws a error", func() {
					achieveIDs = "not valid json"
					achievementIDs, err := client.AchievementIDs()
					Expect(err).To(HaveOccurred())
					Expect(achievementIDs).To(BeNil())
				})
			})
		})

		Context("when requesting 1 or many achievements", func() {
			var (
				statusCode int
			)

			returnedAchievements := []Achievement{
				{
					ID:   1,
					Name: "something",
				},
			}

			BeforeEach(func() {
				statusCode = http.StatusOK
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/"),
						ghttp.RespondWithJSONEncodedPtr(&statusCode, returnedAchievements),
					),
				)
			})

			It("should return the list of returnedAchievements id's", func() {
				achieveIDs := []int{1}

				response, err := client.Achievements(achieveIDs)
				Expect(err).NotTo(HaveOccurred())
				Expect(response).To(Equal(returnedAchievements))
			})
		})
	})

	Describe("DailyCrafting", func() {
		Context("when requesting all daily crafts", func() {
			var statusCode int
			var dailyCrafts []string

			BeforeEach(func() {
				statusCode = http.StatusOK
				dailyCrafts = []string{}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v2/dailycrafting"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &dailyCrafts),
				))
			})

			It("should return the daily crafts", func() {
				dailyCrafts = []string{"craft1", "craft2"}
				resp, err := client.DailyCrafting()
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).To(Equal(dailyCrafts))
			})
		})
	})

	Describe("MapChests", func() {
		Context("when requesting map chests", func() {
			var statusCode int
			var mapchests []string

			BeforeEach(func() {
				statusCode = http.StatusOK
				mapchests = []string{}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v2/mapchests"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &mapchests),
				))
			})

			It("should return the map chests", func() {
				mapchests = []string{"chest1", "chest2"}
				resp, err := client.MapChests()
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).To(Equal(mapchests))
			})
		})
	})

	Describe("WorldBosses", func() {
		Context("when requesting world bosses", func() {
			var statusCode int
			var bosses []string

			BeforeEach(func() {
				statusCode = http.StatusOK
				bosses = []string{}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v2/worldbosses"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &bosses),
				))
			})

			It("should return the world bossses", func() {
				bosses = []string{"chest1", "chest2"}
				resp, err := client.WorldBosses()
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).To(Equal(bosses))
			})
		})
	})
})
