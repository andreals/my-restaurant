package main

import (
	"fmt"
	"log"
	"strings"

	"my-restaurant/pkg/domains/restaurant/model"
	"my-restaurant/pkg/domains/restaurant/repository"
	"my-restaurant/pkg/domains/restaurant/service"
)

func main() {
	var (
		setupRestaurant string
		restaurant      model.Restaurant
		err             error
	)

	svc, err := service.NewService(repository.NewRepository())
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Welcome to application 'My Restaurant'.")
	fmt.Print("Do you want to set up the restaurant? [y/N] ")
	fmt.Scanf("%s", &setupRestaurant)

	if setupRestaurant == "" || strings.ToLower(setupRestaurant) == "n" {
		restaurant, err = svc.LoadRestaurant()
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		restaurant, err = svc.SetupRestaurant()
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	tables, err := svc.TakeOrder(restaurant)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, table := range tables {
		err = svc.PrepareOrder(table, restaurant.Chefs)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
