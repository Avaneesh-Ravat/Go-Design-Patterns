package models

import "github.com/avaneesh-ravat/go-design-pattern/factory"

type ApartmentProperty struct {
	Cost     int
	Layout   string
	Location string
}

func NewApartmentProperty(price int, layout string, location string) *ApartmentProperty {
	return &ApartmentProperty{
		Cost:     price,
		Layout:   layout,
		Location: location,
	}
}

func (h *ApartmentProperty) GetPrice() int {
	return h.Cost
}

func (h *ApartmentProperty) GetLayout() string {
	return h.Layout
}

func (h *ApartmentProperty) GetLocation() string {
	return h.Location
}
