package model

import (
	"fmt"
	"math/big"
)

type ProductKind int

const (
	Apple ProductKind = iota + 1
	Banana
	Pear
	Orange
)

func (pk ProductKind) String() string {
	switch pk {
	case Apple:
		return "apple"
	case Banana:
		return "banana"
	case Pear:
		return "pear"
	case Orange:
		return "orange"
	default:
		return fmt.Sprintf("undefined value: %v", int(pk))
	}
}

type ProductBundle struct {
	Set      map[ProductKind]int
	Discount int
}

type Product struct {
	ID    string
	Kind  ProductKind
	Price *big.Int
}
