package main

import (
	"context"
	"log"
	"time"
)

// Burger Point
const (
	BlueRare = iota
	Rare
	MediumRare
	Medium
	MediumWell
	WellDone
)

// Salad Kind
const (
	WithoutSalad = iota
	Lettuce
	LettuceAndTomatoes
	CreamyColeslaw
	CrispyFriedCabbage
)

// Bacon Kind
const (
	WithoutBacon = iota
	Streaky
	Canadian
	Crumbs
)

type Burger struct {
	State, Point, Salad, Bacon int
	ProductId
	ProductContext
	PreparationTime
}

func RequestBurger(point, salad, bacon int) IProduct {
	return &Burger{
		ProductId: NewProductId(),
		Point:     point,
		Salad:     salad,
		Bacon:     bacon,
		State:     OrderMade,
	}
}

func (p *Burger) GetDefaultPreparationTime() time.Duration {
	return time.Minute * 2
}

func (p *Burger) SendToKitchen() {
	log.Printf("Send Burger {%v} to kitchen", p.GetId().String())

	p.Time = p.GetDefaultPreparationTime()
	switch p.Point {
	case Rare:
		p.Time += time.Second * 30
	case MediumRare:
		p.Time += time.Minute
	case Medium:
		p.Time += time.Minute + (time.Second * 30)
	case MediumWell:
		p.Time += time.Minute * 2
	case WellDone:
		p.Time += (time.Minute * 2) + (time.Second * 30)
	}

	switch p.Bacon {
	case Streaky:
		p.Time += time.Minute * 2
	case Canadian:
		p.Time += time.Minute * 3
	case Crumbs:
		p.Time += time.Minute * 4
	}

	switch p.Salad {
	case Lettuce:
		p.Time += time.Minute
	case LettuceAndTomatoes:
		p.Time += time.Minute + (time.Second + 10)
	case CreamyColeslaw, CrispyFriedCabbage:
		p.Time += time.Minute * 2
	}

	p.State = InProgress
	go p.StartPreparation()
}

func (p *Burger) StartPreparation() {
	p.Ctx, p.Cancel = context.WithTimeout(context.Background(), p.Time)
	select {
	case <-p.Ctx.Done():
		log.Printf("Burger {%s} done", p.GetId().String())
		p.State = Done
	}
}

func (p *Burger) GetState() int {
	return p.State
}

func (p *Burger) GetId() ProductId {
	return p.ProductId
}
