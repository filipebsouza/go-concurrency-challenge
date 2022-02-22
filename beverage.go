package main

import (
	"context"
	"log"
	"time"
)

// Beverage Kind
const (
	Coke = iota
	Soda
	Orange
	Grape
	IcedTea
	OrangeJuice
)

type Beverage struct {
	Kind, State int
	ProductId
	ProductContext
	PreparationTime
}

func RequestBeverage(kind int) IProduct {
	return &Beverage{
		ProductId: NewProductId(),
		Kind:      kind,
		State:     OrderMade,
	}
}

func (p *Beverage) GetDefaultPreparationTime() time.Duration {
	return time.Second * 15
}

func (p *Beverage) SendToKitchen() {
	log.Println("Send Beverage to kitchen")

	preparationTime := p.GetDefaultPreparationTime()
	switch p.Kind {
	case Coke, Soda, Orange, Grape:
		p.Time = preparationTime
	case IcedTea:
		p.Time = preparationTime + (time.Second * 10)
	case OrangeJuice:
		p.Time = preparationTime + (time.Second * 15)
	}

	p.State = InProgress
	go p.StartPreparation()
}

func (p *Beverage) StartPreparation() {
	p.Ctx, p.Cancel = context.WithTimeout(context.Background(), p.Time)
	select {
	case <-p.Ctx.Done():
		log.Println("Beverage done")
		p.State = Done
	}
}

func (p *Beverage) GetState() int {
	return p.State
}

func (p *Beverage) GetId() ProductId {
	return p.ProductId
}
