package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	orderRequests := rand.Int31n(30)
	log.Printf("Number of orders requested are %v", orderRequests)
	for i := int32(0); i < orderRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			numberOfBurgers := rand.Int31n(10)
			numberOfFries := rand.Int31n(10)
			numberOfFries = 0
			numberOfBeverages := rand.Int31n(10)
			numberOfBeverages = 0
			products := make([]IProduct, numberOfFries+numberOfFries+numberOfBeverages)
			for j := int32(0); j < numberOfBurgers; j++ {
				point := rand.Intn(WellDone-BlueRare) + BlueRare
				salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
				bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

				products[j] = RequestBurger(point, salad, bacon)
			}

			order := MakeAnOrder(products...)
			order.WaitFinishPreparation()
		}()
	}

	wg.Wait()
}
