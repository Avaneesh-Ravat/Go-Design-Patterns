package models

type CommercialProperty struct {
	Cost     int
	Layout   string
	Location string
}

func NewCommercialProperty(price int, layout string, location string) *CommercialProperty {
	return &CommercialProperty{
		Cost:     price,
		Layout:   layout,
		Location: location,
	}
}

func (h *CommercialProperty) GetPrice() int {
	return h.Cost
}

func (h *CommercialProperty) GetLayout() string {
	return h.Layout
}

func (h *CommercialProperty) GetLocation() string {
	return h.Location
}
