package model

import (
	"github.com/satori/go.uuid"
	"math/big"
)

const PriceScaleFactor = 100

type Order struct {
	Cart map[ProductKind]OrderProductItem

	DiscountLog []string
	Discounts   []Discount
}

type OrderItem interface {
	Count() int
	Price() *big.Int
}

type OrderProductItem struct {
	Product *Product
	count   int
}

func (o OrderProductItem) Name() string {
	return o.Product.Kind.String()
}

func (o OrderProductItem) Price() *big.Int {
	return o.Product.Price
}

func (o OrderProductItem) Count() int {
	return o.count
}

type OrderSetItem struct {
	SetID string
	count int
	price *big.Int
}

func (o OrderSetItem) Price() *big.Int {
	bigCount := new(big.Int).SetInt64(int64(o.count))
	return o.price.Mul(o.price, bigCount)
}

func (o OrderSetItem) Count() int {
	return o.count
}

func NewOrder() *Order {
	productSetUUID := uuid.NewV4()
	o := &Order{
		Cart: make(map[ProductKind]OrderProductItem),
		Discounts: []Discount{
			&ProductCountDiscount{
				Kind:     Apple,
				MinCount: 7,
				Discount: 10,
			},
			&ProductSetDiscount{
				ID: productSetUUID.String(),
				Set: map[ProductKind]int{
					Banana: 2,
					Pear:   4,
				},
				Discount: 30,
			},
		},
	}

	return o
}

func (o *Order) AddItem(inputProduct *Product, count int) {
	if _, ok := o.Cart[inputProduct.Kind]; ok {
		orderItem := o.Cart[inputProduct.Kind]
		orderItem.count += count
		o.Cart[inputProduct.Kind] = orderItem
		return
	}

	o.Cart[inputProduct.Kind] = OrderProductItem{
		Product: inputProduct,
		count:   count,
	}
}

func (o *Order) Total() float64 {
	total := new(big.Int)

	items := []OrderItem{}
	for _, discount := range o.Discounts {
		orderItem := discount.Modify(o)
		if orderItem != nil {
			items = append(items, orderItem)
		}
	}

	for _, item := range o.Cart {
		if item.count > 0 {
			items = append(items, &OrderProductItem{
				count:   item.count,
				Product: item.Product,
			})
		}
	}

	for _, item := range items {
		price := item.Price()
		price = new(big.Int).SetInt64(1).Mul(price, new(big.Int).SetInt64(int64(item.Count())))
		total.Add(total, price)
	}

	for _, discount := range o.Discounts {
		total.Sub(total, discount.Calculate(items))
	}

	return float64(total.Uint64()) / PriceScaleFactor
}

func (o *Order) AddDiscount(discount Discount) {
	o.Discounts = append(o.Discounts, discount)
}
