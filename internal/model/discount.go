package model

import (
	"math"
	"math/big"
)

type Discount interface {
	Calculate(items []OrderItem) (discount *big.Int)
	Modify(o *Order) (orderItems OrderItem)
}

type ProductCountDiscount struct {
	Kind     ProductKind
	MinCount int
	Discount int
}

func (pcd ProductCountDiscount) Modify(order *Order) OrderItem {
	return nil
}

func (pcd ProductCountDiscount) Calculate(items []OrderItem) *big.Int {
	for _, item := range items {
		switch item.(type) {
		case *OrderProductItem:
			product := item.(*OrderProductItem).Product
			if product.Kind == pcd.Kind && item.Count() >= pcd.MinCount {
				count := new(big.Int).SetInt64(int64(item.Count()))
				totalPrice := new(big.Int).SetInt64(1).Mul(item.Price(), count)
				discount := new(big.Int).SetInt64(int64(pcd.Discount))
				totalPrice.Mul(totalPrice, discount)

				totalPrice.Div(totalPrice, new(big.Int).SetInt64(100))
				return totalPrice
			}
		}
	}

	return new(big.Int).SetInt64(int64(0))
}

type ProductSetDiscount struct {
	ID       string
	Set      map[ProductKind]int
	Discount int
}

func (psd *ProductSetDiscount) Modify(o *Order) OrderItem {
	minSetCount := math.MaxInt32

	pricePerOne := big.NewInt(0)
	for kind, productInSet := range psd.Set {
		for productKind, item := range o.Cart {
			if productKind == kind {
				pricePerOne.Add(pricePerOne, item.Price().Mul(item.Price(), new(big.Int).SetInt64(int64(productInSet))))

				if productInSet > item.count {
					return nil
				} else {
					if minSetCount*productInSet > item.count {
						minSetCount = int(item.count / productInSet)
					}
				}
			}
		}
	}

	if minSetCount == math.MaxInt32 {
		return nil
	}

	// update quantity in cart
	for kind, productInSet := range psd.Set {
		productItem := o.Cart[kind]
		productItem.count -= productInSet * minSetCount
		o.Cart[kind] = productItem
	}

	return &OrderSetItem{
		SetID: psd.ID,
		count: minSetCount,
		price: pricePerOne,
	}
}

func (psd *ProductSetDiscount) Calculate(items []OrderItem) *big.Int {
	for _, item := range items {
		switch item.(type) {
		case *OrderSetItem:
			set := item.(*OrderSetItem)
			if set.SetID == psd.ID {
				totalPrice := new(big.Int).SetInt64(1).Mul(set.price, new(big.Int).SetInt64(int64(set.count*psd.Discount)))
				totalPrice.Div(totalPrice, new(big.Int).SetInt64(100))
				return totalPrice
			}
		}
	}

	return new(big.Int).SetInt64(int64(0))
}

type CouponDiscount struct {
	Kind     ProductKind
	MinCount int
	Discount float64
}

func (cd CouponDiscount) Modify(order *Order) OrderItem {
	return nil
}

func (cd CouponDiscount) Calculate(items []OrderItem) *big.Int {
	for _, item := range items {
		switch item.(type) {
		case *OrderProductItem:
			product := item.(*OrderProductItem).Product
			if product.Kind == cd.Kind {
				count := new(big.Int).SetInt64(int64(item.Count()))
				totalPrice := new(big.Int).SetInt64(1).Mul(item.Price(), count)
				discount := new(big.Int).SetInt64(int64(cd.Discount))
				totalPrice.Mul(totalPrice, discount)

				totalPrice.Div(totalPrice, new(big.Int).SetInt64(100))
				return totalPrice
			}
		}
	}

	return new(big.Int).SetInt64(int64(0))
}
