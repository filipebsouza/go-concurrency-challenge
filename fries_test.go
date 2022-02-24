package main_test

import (
	"github.com/filipebsouza/concurrency-challenge"
	"testing"
	"time"
)

func TestRequestFries_ShouldReturn_ValidFries(t *testing.T) {
	b := main.RequestFries(main.RusticFries)
	if b == nil {
		t.Error("Fries is null")
	}
	if b.GetState() != main.OrderMade {
		t.Error("Fries State is different than OrderMade")
	}
}

func TestFries_GetDefaultPreparationTime_ShouldBe_8Minutes(t *testing.T) {
	b := main.RequestFries(main.Normal)
	if b.GetDefaultPreparationTime() != time.Minute*8 {
		t.Error("Default preparation time for Fries must be 8 minutes")
	}
}

func TestFries_SendToKitchen_StateShouldBe_InProgress(t *testing.T) {
	b := main.RequestFries(main.Spice)
	b.SendToKitchen()
	if b.GetState() != main.InProgress {
		t.Error("Fries order was sent to Kitchen but the State ain't InProgress")
	}
}

func TestFries_StartPreparation_StateShouldBe_Done_AfterPreparation(t *testing.T) {
	b := main.RequestFries(main.Normal)
	b.SendToKitchen()
	duration := b.GetDefaultPreparationTime()
	for {
		if duration <= 0 {
			t.Errorf("Fries preparation exceeded the expected time of %v sec",
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
