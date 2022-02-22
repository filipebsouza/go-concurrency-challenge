package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	orderRequests := rand.Intn(30)
	log.Printf("Number of orders requested are %v", orderRequests)
	wg.Add(orderRequests)
	for i := 0; i < orderRequests; i++ {
		go func() {
			defer wg.Done()
			numberOfBurgers := rand.Intn(10)
			log.Printf("Number of Burgers are %v", numberOfBurgers)
			numberOfFries := rand.Intn(10)
			log.Printf("Number of Fries are %v", numberOfFries)
			numberOfBeverages := rand.Intn(10)
			log.Printf("Number of Beverages are %v", numberOfBeverages)
			products := make([]IProduct, numberOfFries+numberOfFries+numberOfBeverages)

			for j := 0; j < numberOfBurgers; j++ {
				point := rand.Intn(WellDone-BlueRare) + BlueRare
				salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
				bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

				burger := RequestBurger(point, salad, bacon)
				products = append(products, burger)
			}

			for j := 0; j < numberOfFries; j++ {
				kind := rand.Intn(RusticFries-Normal) + Normal

				fries := RequestFries(kind)
				products = append(products, fries)
			}

			for j := 0; j < numberOfBeverages; j++ {
				kind := rand.Intn(OrangeJuice-Coke) + Coke

				beverage := RequestBeverage(kind)
				products = append(products, beverage)
			}

			order := MakeAnOrder(products...)
			order.WaitFinishPreparation()
		}()
	}

	wg.Wait()
}
