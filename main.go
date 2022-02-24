package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	orderRequests := 0
	for {
		orderRequests = rand.Intn(30)
		if orderRequests > 0 {
			break
		}
	}

	log.Printf("Number of orders requested are %v", orderRequests)
	wg.Add(orderRequests)
	for i := 0; i < orderRequests; i++ {
		go func() {
			defer wg.Done()
			numberOfBurgers, numberOfFries, numberOfBeverages, quantity := 0, 0, 0, 0
			for {
				numberOfBurgers = rand.Intn(10)
				log.Printf("Number of Burgers are %v", numberOfBurgers)
				numberOfFries = rand.Intn(10)
				log.Printf("Number of Fries are %v", numberOfFries)
				numberOfBeverages = rand.Intn(10)
				log.Printf("Number of Beverages are %v", numberOfBeverages)
				quantity = numberOfBurgers + numberOfFries + numberOfBeverages

				if quantity > 0 {
					break
				}
			}
			products := make([]IProduct, quantity)
			index := 0
			for {
				for j := 0; j < numberOfBurgers; j++ {
					point := rand.Intn(WellDone-BlueRare) + BlueRare
					salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
					bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

					burger := RequestBurger(point, salad, bacon)
					products[index] = burger
					index++
				}

				for j := 0; j < numberOfFries; j++ {
					kind := rand.Intn(RusticFries-Normal) + Normal

					fries := RequestFries(kind)
					products[index] = fries
					index++
				}

				for j := 0; j < numberOfBeverages; j++ {
					kind := rand.Intn(OrangeJuice-Coke) + Coke

					beverage := RequestBeverage(kind)
					products[index] = beverage
					index++
				}

				if index == quantity {
					break
				}
			}

			order := MakeAnOrder(products...)
			order.WaitFinishPreparation()
		}()
	}

	wg.Wait()
	log.Println("All orders finished")
}

func requestProducts() []IProduct {
	numberOfBurgers, numberOfFries, numberOfBeverages, quantity := 0, 0, 0, 0
	for {
		numberOfBurgers = rand.Intn(10)
		log.Printf("Number of Burgers are %v", numberOfBurgers)
		numberOfFries = rand.Intn(10)
		log.Printf("Number of Fries are %v", numberOfFries)
		numberOfBeverages = rand.Intn(10)
		log.Printf("Number of Beverages are %v", numberOfBeverages)
		quantity = numberOfBurgers + numberOfFries + numberOfBeverages

		if quantity > 0 {
			break
		}
	}
	products := make([]IProduct, quantity)
	index := 0
	for {
		var wg sync.WaitGroup
		wg.Add(3)
		burgers := make([]IProduct, numberOfBurgers)
		go requestBurgers(wg, numberOfBurgers, &burgers)

		for j := 0; j < numberOfBurgers; j++ {
			point := rand.Intn(WellDone-BlueRare) + BlueRare
			salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
			bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

			burger := RequestBurger(point, salad, bacon)
			products[index] = burger
			index++
		}

		for j := 0; j < numberOfFries; j++ {
			kind := rand.Intn(RusticFries-Normal) + Normal

			fries := RequestFries(kind)
			products[index] = fries
			index++
		}

		for j := 0; j < numberOfBeverages; j++ {
			kind := rand.Intn(OrangeJuice-Coke) + Coke

			beverage := RequestBeverage(kind)
			products[index] = beverage
			index++
		}

		if index == quantity {
			break
		}
	}

	return products
}

func requestBurgers(wg sync.WaitGroup, numberOfBurgers int, burgers *[]IProduct) interface{} {
	defer wg.Done()
	for i := 0; i < numberOfBurgers; i++ {
		point := rand.Intn(WellDone-BlueRare) + BlueRare
		salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
		bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

		burger := RequestBurger(point, salad, bacon)
		(*burgers)[i] = burger
	}

	return burgers
}
