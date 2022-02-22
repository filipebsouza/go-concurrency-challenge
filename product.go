package main

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Product State
const (
	OrderMade = iota
	InProgress
	Done
)

type IProduct interface {
	SendToKitchen()
	StartPreparation()
	GetDefaultPreparationTime() time.Duration
	GetState() int
	GetId() ProductId
}

type PreparationTime struct {
	Time, Elapsed time.Duration
}

type ProductContext struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

type ProductId uuid.UUID

func (productId ProductId) String() string {
	return uuid.UUID(productId).String()
}

func NewProductId() ProductId {
	return ProductId(uuid.New())
}
