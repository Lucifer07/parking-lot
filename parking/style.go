package parking

import (
	"sort"
)

type StyleContract interface {
	GetParkingLot(lots []*ParkingLot) *ParkingLot
}
type DefaultStyle struct{}

func StyleDefault() *DefaultStyle {
	return &DefaultStyle{}
}

func (d *DefaultStyle) GetParkingLot(lots []*ParkingLot) *ParkingLot {
	return lots[0]
}

type SortingCapacity struct{}

func StyleCapacity() *SortingCapacity {
	return &SortingCapacity{}
}
func (c *SortingCapacity) GetParkingLot(lots []*ParkingLot) *ParkingLot {
	sort.Slice(lots, func(i, j int) bool {
		return lots[i].max > lots[j].max
	})
	return lots[0]
}

type SortingFreeSpace struct{}

func StyleFreeSpace() *SortingFreeSpace {
	return &SortingFreeSpace{}
}
func (f *SortingFreeSpace) GetParkingLot(lots []*ParkingLot) *ParkingLot {
	sort.Slice(lots, func(i, j int) bool {
		return lots[i].max-len(lots) > lots[j].max-len(lots)
	})
	return lots[0]
}
