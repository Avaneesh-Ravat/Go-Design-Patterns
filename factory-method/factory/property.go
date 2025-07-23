package factory

import (
	"fmt"
	"github.com/avaneesh-ravat/go-design-patterns/factory-method/models"
)

type Property interface {
	GetPrice() int
	GetLayout() string
	GetLocation() string
}

func CreateProperty(propertyType string, price int, layout string, location string) (Property, error) {
	switch propertyType {
	case HouseProperty:
		return models.NewHouseProperty(price, layout, location), nil
	case ApartmentProperty:
		return models.NewApartmentProperty(price, layout, location), nil
	case CommercialProperty:
		return models.NewCommercialProperty(price, layout, location), nil
	default:
		return nil, fmt.Errorf("Invalid property type: %s", propertyType)
	}
}
