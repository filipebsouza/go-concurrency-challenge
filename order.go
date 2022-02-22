package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

const (
	ValidationProductIsRequired = "An order must have at least one product"
)

const (
	Requested = iota
	Making
	ReadyForDelivery
)

type OrderId uuid.UUID

type IOrder interface {
	GetOrderId() OrderId
	WaitFinishPreparation()
	GetProducts() *[]IProduct
	GetOrderStatus() int
}

type order struct {
	id       OrderId
	products *[]IProduct
	status   int
}

func MakeAnOrder(products ...IProduct) IOrder {
	log.Println("Make an order")

	if len(products) <= 0 {
		log.Fatal(ValidationProductIsRequired)
	}

	o := &order{
		OrderId(uuid.New()),
		&products,
		Requested,
	}

	for _, product := range *o.products {
		go product.SendToKitchen()
	}

	return o
}

func (o *order) GetOrderId() OrderId {
	return o.id
}

func (o *order) WaitFinishPreparation() {
	log.Println("Wait the preparation")
	o.status = Making
	productsToVerify := *o.products
	length := len(productsToVerify)
	for {
		for i := 0; i < length; i++ {
			product := productsToVerify[i]
			if product.GetState() == Done {
				productsToVerify = removeProduct(productsToVerify, product.GetId())
				break
			}
		}

		length = len(productsToVerify)
		if length <= 0 {
			break
		}

		time.Sleep(time.Second * 30)
	}

	o.status = ReadyForDelivery
	log.Println("Order finish")
}

func (o *order) GetProducts() *[]IProduct {
	return o.products
}

func (o *order) GetOrderStatus() int {
	return o.status
}

func removeProduct(products []IProduct, idToRemove ProductId) []IProduct {
	length := len(products)
	for i, product := range products {
		if idToRemove == product.GetId() {
			products[length-1], products[i] = products[i], products[length-1]
			return products[:length-1]
		}
	}
	return products
}
