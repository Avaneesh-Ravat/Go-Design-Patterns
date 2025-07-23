package models

type HouseProperty struct {
	Cost     int
	Layout   string
	Location string
}

func NewHouseProperty(price int, layout string, location string) *HouseProperty {
	return &HouseProperty{
		Cost:     price,
		Layout:   layout,
		Location: location,
	}
}

func (h *HouseProperty) GetPrice() int {
	return h.Cost
}

func (h *HouseProperty) GetLayout() string {
	return h.Layout
}

func (h *HouseProperty) GetLocation() string {
	return h.Location
}
