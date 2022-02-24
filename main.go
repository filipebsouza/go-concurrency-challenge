package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	orderRequests := 0
	for {
		orderRequests = rand.Intn(30)
		if orderRequests > 0 {
			break
		}
	}

	log.Printf("Number of orders requested are %v", orderRequests)
	wg := new(sync.WaitGroup)
	wg.Add(orderRequests)
	for i := 0; i < orderRequests; i++ {
		go func() {
			defer wg.Done()

			wgProduct := new(sync.WaitGroup)
			cn := make(chan []IProduct, 1)
			wgProduct.Add(1)

			go requestProducts(wgProduct, cn)

			wgProduct.Wait()

			products := <-cn
			close(cn)
			order := MakeAnOrder(products...)
			order.WaitFinishPreparation()
		}()
	}

	wg.Wait()
	log.Println("All orders finished")
}

func requestProducts(wg *sync.WaitGroup, result chan []IProduct) {
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
	wgProducts := new(sync.WaitGroup)
	wgProducts.Add(3)
	burgers := make([]IProduct, numberOfBurgers)
	fries := make([]IProduct, numberOfFries)
	beverages := make([]IProduct, numberOfBeverages)

	go requestBurgers(wgProducts, numberOfBurgers, &burgers)
	go requestFries(wgProducts, numberOfFries, &fries)
	go requestBeverages(wgProducts, numberOfBeverages, &beverages)

	wgProducts.Wait()

	products := burgers
	products = append(products, fries[:]...)
	products = append(products, beverages[:]...)

	result <- products
}

func requestBurgers(wg *sync.WaitGroup, numberOfBurgers int, burgers *[]IProduct) {
	defer wg.Done()
	for i := 0; i < numberOfBurgers; i++ {
		point := rand.Intn(WellDone-BlueRare) + BlueRare
		salad := rand.Intn(CrispyFriedCabbage-WithoutSalad) + WithoutSalad
		bacon := rand.Intn(Crumbs-WithoutBacon) + WithoutBacon

		burger := RequestBurger(point, salad, bacon)
		(*burgers)[i] = burger
	}
}

func requestFries(wg *sync.WaitGroup, numberOfFries int, fries *[]IProduct) {
	defer wg.Done()
	for i := 0; i < numberOfFries; i++ {
		kind := rand.Intn(RusticFries-Normal) + Normal

		f := RequestFries(kind)
		(*fries)[i] = f
	}
}

func requestBeverages(wg *sync.WaitGroup, numberOfBeverages int, beverages *[]IProduct) {
	defer wg.Done()
	for i := 0; i < numberOfBeverages; i++ {
		kind := rand.Intn(OrangeJuice-Coke) + Coke

		beverage := RequestBeverage(kind)
		(*beverages)[i] = beverage
	}
}
