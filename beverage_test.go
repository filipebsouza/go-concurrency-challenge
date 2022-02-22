package main_test

import (
	"github.com/filipebsouza/concurrency-challenge"
	"testing"
	"time"
)

func TestRequestBeverage_ShouldReturn_ValidBeverage(t *testing.T) {
	b := main.RequestBeverage(main.Coke)
	if b == nil {
		t.Error("Beverage is null")
	}
	if b.GetState() != main.OrderMade {
		t.Error("Beverage State is different than OderMade")
	}
}

func TestBeverage_GetDefaultPreparationTime_ShouldBe_15Seconds(t *testing.T) {
	b := main.RequestBeverage(main.Coke)
	if b.GetDefaultPreparationTime() != time.Second*15 {
		t.Error("Default preparation time for Beverage must be 15 seconds")
	}
}

func TestBeverage_SendToKitchen_StateShouldBe_InProgress(t *testing.T) {
	b := main.RequestBeverage(main.Coke)
	b.SendToKitchen()
	if b.GetState() != main.InProgress {
		t.Error("Beverage order was sent to Kitchen but the State ain't InProgress")
	}
}

func TestBeverage_StartPreparation_StateShouldBe_Done_AfterPreparation(t *testing.T) {
	b := main.RequestBeverage(main.IcedTea)
	b.SendToKitchen()
	icedTeaIncreasedTime := time.Second * 10
	duration := b.GetDefaultPreparationTime()
	duration += icedTeaIncreasedTime
	for {
		if duration <= 0 {
			t.Errorf("Beverage preparation exceeded the expected time of %v sec",
				b.GetDefaultPreparationTime()+icedTeaIncreasedTime)
			break
		}
		if b.GetState() == main.Done {
			break
		}
		time.Sleep(time.Second)
		duration--
	}
}
