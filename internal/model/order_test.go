package model_test

import (
	"github.com/mchusovlianov/cinemo/internal/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order test")
}

var _ = BeforeSuite(func() {
})

var _ = Describe("Order test", func() {
	var ()
	Context("Error configuration", func() {
		It("Check simple calculation", func() {
			a := &model.Product{
				Kind:  model.Apple,
				Price: new(big.Int).SetInt64(2 * model.PriceScaleFactor),
			}

			order := model.NewOrder()
			order.AddItem(a, 3)
			total := order.Total()
			Expect(total).Should(Equal(6.0))
		})

		It("Check product discount calculation", func() {
			order := model.NewOrder()
			a := &model.Product{
				Kind:  model.Apple,
				Price: new(big.Int).SetInt64(2 * model.PriceScaleFactor),
			}

			// If 7 or more apples are added to the cart, a 10% discount is applied to all apples.
			order.AddItem(a, 3)
			order.AddItem(a, 4)
			total := order.Total()
			Expect(total).Should(Equal(12.6))
		})

		It("Check set discount", func() {
			//For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set. These sets must be added to their own cart item entry.
			order := model.NewOrder()
			pear := &model.Product{
				Kind:  model.Pear,
				Price: new(big.Int).SetInt64(7 * model.PriceScaleFactor),
			}

			banana := &model.Product{
				Kind:  model.Banana,
				Price: new(big.Int).SetInt64(9 * model.PriceScaleFactor),
			}

			order.AddItem(pear, 4)
			order.AddItem(banana, 2)

			total := order.Total()
			Expect(total).Should(Equal(32.2))
		})

		It("Check coupon discount", func() {
			//For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set. These sets must be added to their own cart item entry.
			order := model.NewOrder()
			order.AddDiscount(&model.CouponDiscount{
				Kind:     model.Orange,
				Discount: 30,
			})

			orange := &model.Product{
				Kind:  model.Orange,
				Price: new(big.Int).SetInt64(7 * model.PriceScaleFactor),
			}

			order.AddItem(orange, 10)

			total := order.Total()
			Expect(total).Should(Equal(49.0))
		})
	})
})

var _ = AfterSuite(func() {})
