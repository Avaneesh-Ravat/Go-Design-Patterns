package factorymethod

import (
	"fmt"

	"github.com/avaneesh-ravat/go-design-patterns/factory-method/factory"
)

func FactoryMethod() {
	fmt.Println("Welcome to design patterns repo, we will be learning the factory-method pattern here.")
	fmt.Printf("\nHere, we have a example of real state properties, we have 3 kind of properties.\nIn this code the main function will not be creating the objects of each property type it will be handled by factory\nWe need to just pass the some values to create the property data.")

	house, _ := factory.CreateProperty("house", 10000000, "3BHK", "Noida")
	apartment, _ := factory.CreateProperty("apartment", 1500000, "2BHK", "Hyderabad")
	commercial, _ := factory.CreateProperty("commercial", 2000000, "40*60", "Gurugram")

	fmt.Printf("House Property details: \nCost: %d\nLayout: %s\nLocation:%s", house.GetPrice(), house.GetLayout(), house.GetLocation())

}
