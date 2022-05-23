package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-restaurant/pkg/domains/restaurant/model"
	"os"
	"strings"
	"sync"
	"time"
)

// RepositoryI is a interface to communicate with a external source of data
type RepositoryI interface {
	Execer
}

type Execer interface {
	LoadRestaurant() (model.Restaurant, error)
	SetupRestaurant() (model.Restaurant, error)
	TakeOrder(model.Restaurant) ([]model.Table, error)
	PrepareOrder(model.Table, int64) error
}

type RepositoryMemory struct{}

// NewRepository generates a struct of type RepositoryMemory
func NewRepository() *RepositoryMemory {
	return &RepositoryMemory{}
}

func (r *RepositoryMemory) LoadRestaurant() (model.Restaurant, error) {
	var (
		indexRestaurant int
		restaurant      model.Restaurant
	)

	fmt.Println("Do you want to load which restaurant?")
	files, err := ioutil.ReadDir("restaurants/")
	if err != nil {
		return restaurant, err
	}

	for idx, file := range files {
		fmt.Printf("[%d] %s\n", idx, strings.Replace(file.Name(), ".json", "", -1))
	}

	fmt.Print("Enter the number of restaurant: ")
	fmt.Scanf("%d", &indexRestaurant)
	if indexRestaurant < 0 || indexRestaurant > len(files) {
		return restaurant, fmt.Errorf("invalid value")
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("restaurants/%s", files[indexRestaurant].Name()))
	if err != nil {
		return restaurant, err
	}

	err = json.Unmarshal(content, &restaurant)
	if err != nil {
		return restaurant, err
	}

	fmt.Print("Restaurant successfully loaded\n\n")
	return restaurant, nil
}

func (r *RepositoryMemory) SetupRestaurant() (model.Restaurant, error) {
	var (
		restaurant model.Restaurant
		err        error
	)

	fmt.Print("Enter the name of the restaurant: ")
	restName, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return restaurant, err
	}

	restaurant.Name = strings.Replace(restName, "\n", "", -1)
	fmt.Print("Enter the number of chefs the restaurant has: ")
	fmt.Scanf("%d", &restaurant.Chefs)

	fmt.Println("Let's register the restaurant menu.")
	fmt.Println("To stop registering, just type 'exit'")
	for {
		var (
			dish        model.Dish
			preparation int64
			err         error
		)

		fmt.Print("Enter the dish name: ")
		dishName, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return restaurant, err
		}

		dish.Name = strings.Replace(dishName, "\n", "", -1)
		if strings.ToLower(dish.Name) == "exit" {
			break
		}

		fmt.Print("Enter the price: ")
		fmt.Scanf("%f", &dish.Price)

		fmt.Print("Enter preparation time (in seconds): ")
		fmt.Scanf("%d", &preparation)

		dish.PreparationTime, err = time.ParseDuration(fmt.Sprintf("%ds", preparation))
		if err != nil {
			return restaurant, err
		}
		restaurant.Menu.Dishes = append(restaurant.Menu.Dishes, dish)
	}

	content, err := json.Marshal(restaurant)
	if err != nil {
		return restaurant, err
	}

	err = ioutil.WriteFile(fmt.Sprintf("restaurants/%s.json", restaurant.Name), content, 0644)
	if err != nil {
		return restaurant, err
	}

	fmt.Print("Restaurant successfully configured\n\n")
	return restaurant, nil
}

func (r *RepositoryMemory) TakeOrder(rest model.Restaurant) ([]model.Table, error) {
	var tables []model.Table
	fmt.Println("Let's take orders!")

	for {
		var (
			table  model.Table
			orders []model.Order
		)

		if len(tables) > 0 {
			var anotherTable string
			fmt.Print("Take order for another table? [y/N] ")
			fmt.Scanf("%s", &anotherTable)

			if anotherTable == "" || strings.ToLower(anotherTable) == "n" {
				break
			}
		}

		fmt.Print("Enter the table number: ")
		fmt.Scanf("%d", &table.Number)

		fmt.Print("Enter the number of customers: ")
		fmt.Scanf("%d", &table.Customers)

		for i := 0; i < int(table.Customers); i++ {
			var (
				order  model.Order
				dishes []model.Dish
			)
			fmt.Println("\n========================================================================================")
			fmt.Printf("Restaurant %s - Register the order for Customer %d.\n", rest.Name, i+1)
			fmt.Println("======================================== Menu ==========================================")
			for idx, dish := range rest.Menu.Dishes {
				fmt.Printf("[%d] %s - Price: %.2f - Preparation Time: %s\n", idx+1, dish.Name, dish.Price, dish.PreparationTime)
			}
			fmt.Println("========================================================================================")
			fmt.Printf("\nTo stop registering order for customer %d, just type '-1'\n", i+1)
			for {
				var (
					dishNumber int
					dish       model.Dish
				)

				fmt.Print("Enter the dish number: ")
				fmt.Scanf("%d", &dishNumber)

				if dishNumber == -1 {
					break
				}

				if dishNumber-1 < 0 || dishNumber-1 > len(rest.Menu.Dishes) {
					return tables, fmt.Errorf("invalid value")
				}

				dish = rest.Menu.Dishes[dishNumber-1]
				dishes = append(dishes, dish)
			}

			order.Customer = int64(i + 1)
			order.Dishes = dishes

			orders = append(orders, order)
		}

		table.Orders = orders
		tables = append(tables, table)
	}

	fmt.Println("Order sent to kitchen")
	return tables, nil
}

func (r *RepositoryMemory) PrepareOrder(table model.Table, chefs int64) error {
	var (
		wg = sync.WaitGroup{}
	)

	orders, err := splitOrders(table.Orders, int(chefs))
	if err != nil {
		return err
	}

	fmt.Printf("The kitchen received the order for table #%d\n", table.Number)

	for chef := int64(0); chef < chefs; chef++ {
		wg.Add(1)

		go func(orders []model.Order, chef int64) {
			defer wg.Done()
			for _, order := range orders {

				var sleepTime time.Duration
				fmt.Printf("Chef %d is preparing the order for customer %d\n", chef+1, order.Customer)
				for _, dish := range order.Dishes {
					if dish.PreparationTime.Seconds() > sleepTime.Seconds() {
						sleepTime = dish.PreparationTime
					}
				}

				fmt.Printf("Wait %s while the chef %d prepares the dishes.\n", sleepTime, chef+1)
				time.Sleep(sleepTime)
			}
		}(orders[chef], chef)
	}

	wg.Wait()
	fmt.Printf("Table #%d order completed. Serving customers...\n", table.Number)
	return generateBill(table)
}

func generateBill(table model.Table) error {
	var totalValue float64
	fmt.Printf("========== Bill of Table #%d ==========\n", table.Number)
	for _, order := range table.Orders {
		var totalCustomer float64
		fmt.Printf("- Customer %d\n", order.Customer)
		for _, dish := range order.Dishes {
			fmt.Printf("	%s - Price: $ %.2f\n", dish.Name, dish.Price)
			totalCustomer += dish.Price
			totalValue += dish.Price
		}
		fmt.Printf("- Sub-Total: $ %.2f\n\n", totalCustomer)
	}
	fmt.Printf("	Total: $ %.2f\n", totalValue)
	fmt.Println("=======================================")
	return nil
}

func splitOrders(orders []model.Order, chefs int) ([][]model.Order, error) {
	var (
		chunkSize = (len(orders) + chefs - 1) / chefs
		divided   = [][]model.Order{}
	)

	for i := 0; i < len(orders); i += chunkSize {
		end := i + chunkSize

		if end > len(orders) {
			end = len(orders)
		}

		divided = append(divided, orders[i:end])
	}
	return divided, nil
}
