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

func (orderId OrderId) String() string {
	return uuid.UUID(orderId).String()
}

type IOrder interface {
	GetId() OrderId
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

	log.Printf("Sending products of order {%s} to kitchen", o.GetId().String())
	for _, product := range *o.products {
		go product.SendToKitchen()
	}

	return o
}

func (o *order) GetId() OrderId {
	return o.id
}

func (o *order) WaitFinishPreparation() {
	log.Printf("Wait the preparation of order {%s}", o.GetId().String())
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

		log.Printf("Still waiting for the order {%s}", o.GetId().String())
		time.Sleep(time.Second * 10)
	}

	o.status = ReadyForDelivery
	log.Printf("Order {%s} finish", o.GetId().String())
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
