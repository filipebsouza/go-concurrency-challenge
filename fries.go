package main

import (
	"context"
	"log"
	"time"
)

const (
	Normal = iota
	Spice
	RusticFries
)

type Fries struct {
	Kind, State int
	ProductId
	ProductContext
	PreparationTime
}

func RequestFries(kind int) IProduct {
	return &Fries{
		ProductId: NewProductId(),
		Kind:      kind,
		State:     OrderMade,
	}
}

func (p *Fries) GetDefaultPreparationTime() time.Duration {
	return time.Minute * 8
}

func (p *Fries) SendToKitchen() {
	log.Printf("Send Fries {%s} to kitchen", p.GetId().String())

	preparationTime := p.GetDefaultPreparationTime()
	switch p.Kind {
	case Normal:
		p.Time = preparationTime
	case Spice:
		p.Time = preparationTime + (time.Minute * 2)
	case RusticFries:
		p.Time = preparationTime + (time.Minute * 4)
	}

	p.State = InProgress
	go p.StartPreparation()
}

func (p *Fries) StartPreparation() {
	p.Ctx, p.Cancel = context.WithTimeout(context.Background(), p.Time)
	select {
	case <-p.Ctx.Done():
		log.Printf("Fries {%s} done", p.GetId().String())
		p.State = Done
	}
}

func (p *Fries) GetState() int {
	return p.State
}

func (p *Fries) GetId() ProductId {
	return p.ProductId
}
