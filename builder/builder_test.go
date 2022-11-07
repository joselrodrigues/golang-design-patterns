package builder_test

import (
	m "patterns/builder/mock"
	"sort"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	b "patterns/builder"
)

var _ = Describe("Builder", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	When("CookChocalateCake is called", func() {
		Context("AddMixtureToMold", func() {
			It("should return nill if MoldSize is correct", func() {
				defer mockCtrl.Finish()
				builder := b.GetBuilder("ChocolateLarge")
				director := b.SetDirector(builder)
				m := m.NewMockCakeBuilder(mockCtrl)
				m.
					EXPECT().
					AddMixtureToMold().
					Return(nil).
					AnyTimes()

				director.CookChocalateCake()
			})

			It("should panic if MoldSize is not correct", func() {
				actual := func() {
					builder := b.SetCakeBuilder()
					builder.MoldSize = "Whatever"
					builder.Ingredients = []string{"sugar"}
					builder.OvenTemperature = 40
					builder.ShakingMinuts = 1
					director := b.SetDirector(builder)
					director.CookChocalateCake()
				}
				expect := gomega.Panic()
				gomega.Expect(actual).To(expect)

			})

			It("should cake has MoldSize L if builder is ChocolateLarge", func() {
				builder := b.GetBuilder("ChocolateLarge")
				director := b.SetDirector(builder)

				director.CookChocalateCake()

				actual := builder.GetCake().MoldSize
				expect := b.LARGE

				gomega.Expect(actual).To(gomega.Equal(expect))

			})

			It("should cake has MoldSize S if builder is ChocolateSmall", func() {
				builder := b.GetBuilder("ChocolateSmall")
				director := b.SetDirector(builder)

				director.CookChocalateCake()

				actual := builder.GetCake().MoldSize
				expect := b.SMALL

				gomega.Expect(actual).To(gomega.Equal(expect))

			})
		})

		Context("ShakeMixture", func() {
			It("should have 4 ingredients if builder is ChocolateLarge", func() {
				builder := b.GetBuilder("ChocolateLarge")
				director := b.SetDirector(builder)

				director.CookChocalateCake()

				actual := len(builder.GetCake().Ingredients)
				expect := 4

				gomega.Expect(actual).To(gomega.Equal(expect))
			})

			It("should have 4 ingredients if builder is ChocolateLarge", func() {
				builder := b.GetBuilder("ChocolateSmall")
				director := b.SetDirector(builder)

				director.CookChocalateCake()

				actual := len(builder.GetCake().Ingredients)
				expect := 4

				gomega.Expect(actual).To(gomega.Equal(expect))
			})
		})

		Context("CookInOven", func() {
			It("should change fahrenheit to celsius in ChocolateSmall builder", func() {
				builder := b.GetBuilder("ChocolateSmall")
				director := b.SetDirector(builder)
				fahrenheitTemp := builder.GetCake().OvenTemperature

				director.CookChocalateCake()

				actual := builder.GetCake().OvenTemperature
				expect := (fahrenheitTemp - 32) * 5 / 9

				gomega.Expect(actual).To(gomega.Equal(expect))
			})
		})

	})
	When("GetBuilder", func() {
		It("should return nil if argument is not valid", func() {
			builder := b.GetBuilder("invalid")

			actual := builder
			gomega.Expect(actual).To(gomega.BeNil())
		})
	})
	When("MixIngredients is called", func() {
		It("should ingredients be sorted", func() {
			builder := b.SetCakeBuilder()
			builder.Ingredients = []string{"sugar", "butter", "eggs", "chocolate"}

			builder.MixIngredients()

			cake := builder.GetCake()
			sort.StringsAreSorted(cake.Ingredients)

			actual := sort.StringsAreSorted(cake.Ingredients)
			expect := gomega.BeTrue()

			gomega.Expect(actual).To(expect)

		})
	})
})
