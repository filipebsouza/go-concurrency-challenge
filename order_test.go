package main_test

import (
	"github.com/filipebsouza/concurrency-challenge"
	"testing"
)

func TestMakeAnOrder_ShouldReturn_AnError_When_NoProductIsInformed(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()

	t.Error(main.MakeAnOrder(nil))
}

func TestMakeAnOrder_ShouldReturn_ValidOrder(t *testing.T) {
	testCases := []struct {
		name     string
		products []main.IProduct
		error    string
	}{
		{
			"Make order with Burger without salad and without bacon",
			[]main.IProduct{
				main.RequestBurger(main.Medium, main.WithoutSalad, main.WithoutBacon),
			},
			"Expect one Burger, but product length doesn't match",
		},
		{
			"Make order with 2 Burgers",
			[]main.IProduct{
				main.RequestBurger(main.Medium, main.WithoutSalad, main.WithoutBacon),
				main.RequestBurger(main.BlueRare, main.CreamyColeslaw, main.Canadian),
			},
			"Expected 2 Burger, but product length doesn't match",
		},
		{
			"Make order with Burger, Fries and Beverage",
			[]main.IProduct{
				main.RequestBurger(main.WellDone, main.Lettuce, main.Crumbs),
				main.RequestFries(main.Spice),
				main.RequestBeverage(main.IcedTea),
			},
			"Expected a Burger, Fries and Beverage, but product length doesn't match",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			o := main.MakeAnOrder(test.products...)
			if len(*o.GetProducts()) != len(test.products) {
				t.Log(test.error)
			}
		})
	}
}

func TestWaitFinishPreparation_ShouldNotReturn_AnError(t *testing.T) {
	coke := main.RequestBeverage(main.Coke)
	p := &coke
	o := main.MakeAnOrder(*p)
	for {
		o.WaitFinishPreparation()
		if (*p).GetState() == main.Done {
			break
		}
	}

	if o.GetOrderStatus() != main.ReadyForDelivery {
		t.Error("Order unfinished but products was done.")
	}
}
