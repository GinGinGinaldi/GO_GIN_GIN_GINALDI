package main

import "fmt"

type Vehicle struct {
	Brand string
	Type  string
	Rent  int
}

type Customer struct {
	Name           string
	Address        string
	RentedVehicles []Vehicle
}

func (c *Customer) rentVehicle(v Vehicle) {
	c.RentedVehicles = append(c.RentedVehicles, v)
}

func (c *Customer) returnVehicle(index int) {
	if index >= 0 && index < len(c.RentedVehicles) {
		c.RentedVehicles = append(c.RentedVehicles[:index], c.RentedVehicles[index+1:]...)
	}
}

func (c Customer) showRentedVehicles() {
	fmt.Printf("Customer: %s, Address: %s\n", c.Name, c.Address)
	for _, v := range c.RentedVehicles {
		fmt.Printf("- %s %s, Rent: %d\n", v.Brand, v.Type, v.Rent)
	}
}

func main() {
	customer := Customer{Name: "Gin", Address: "Bandung"}
	car := Vehicle{Brand: "Toyota", Type: "Avanza", Rent: 500000}
	bike := Vehicle{Brand: "Honda", Type: "Vario", Rent: 150000}

	customer.rentVehicle(car)
	customer.rentVehicle(bike)

	customer.showRentedVehicles()

	customer.returnVehicle(0)
	fmt.Println("After returning a vehicle:")
	customer.showRentedVehicles()
}
