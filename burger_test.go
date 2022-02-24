package main_test

import (
	"github.com/filipebsouza/concurrency-challenge"
	"testing"
	"time"
)

func TestRequestBurger_ShouldReturn_ValidBurger(t *testing.T) {
	b := main.RequestBurger(main.WellDone, main.CreamyColeslaw, main.Canadian)
	if b == nil {
		t.Error("Burger is null")
	}
	if b.GetState() != main.OrderMade {
		t.Error("Burger State is different than OrderMade")
	}
}

func TestBurger_GetDefaultPreparationTime_ShouldBe_2Minutes(t *testing.T) {
	b := main.RequestBurger(main.BlueRare, main.WithoutSalad, main.WithoutBacon)
	if b.GetDefaultPreparationTime() != time.Minute*2 {
		t.Error("Default preparation time for Burger must be 2 minutes")
	}
}

func TestBurger_SendToKitchen_StateShouldBe_InProgress(t *testing.T) {
	b := main.RequestBurger(main.Medium, main.LettuceAndTomatoes, main.Crumbs)
	b.SendToKitchen()
	if b.GetState() != main.InProgress {
		t.Error("Burger order was sent to Kitchen but the State ain't InProgress")
	}
}

func TestBurger_StartPreparation_StateShouldBe_Done_AfterPreparation(t *testing.T) {
	b := main.RequestBurger(main.BlueRare, main.WithoutSalad, main.WithoutBacon)
	b.SendToKitchen()
	duration := b.GetDefaultPreparationTime()
	for {
		if duration <= 0 {
			t.Errorf("Burger preparation exceeded the expected time of %v sec",
				b.GetDefaultPreparationTime())
			break
		}
		if b.GetState() == main.Done {
			break
		}
		time.Sleep(time.Second)
		duration--
	}
}
