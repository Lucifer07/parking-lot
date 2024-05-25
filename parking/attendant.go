package parking

import (
	"errors"

	"github.com/Lucifer07/parking-lot/entity"
)

var (
	ErrorNilLot      = errors.New("you dont have Parking lot")
	ErrorParkingFull = errors.New("parking all full")
	ErrorLocation    = errors.New("location not found")
)

type Attendant struct {
	area         []*ParkingLot
	availablelot []*ParkingLot
	style        StyleContract
}

type Observer interface {
	UpdateFull(*ParkingLot)
	UpdateAvailable(*ParkingLot)
}

func (a *Attendant) ChangeStyle(s StyleContract) {
	a.style = s
}
func (a *Attendant) UpdateFull(p *ParkingLot) {
	a.availablelot = deleteLot(a.availablelot, p)
}

func (a *Attendant) UpdateAvailable(p *ParkingLot) {
	a.availablelot = append(a.availablelot, p)
}
func (a *Attendant) AddParkinglot(max int) *Attendant {
	lot := NewParkingLot(max)
	lot.Register(a)
	a.area = append(a.area, lot)
	a.availablelot = append(a.availablelot, lot)
	return a
}
func (a *Attendant) Park(car *entity.Car) (*entity.Ticket, error) {
	if car == nil {
		return nil, ErrorNil
	}
	if len(a.availablelot) != 0 {
		lot := a.style.GetParkingLot(a.availablelot)
		if lot != nil {
			return lot.Park(car)
		}

	}
	return nil, ErrorParkingFull
}
func (a *Attendant) Location(ticket *entity.Ticket) (*ParkingLot, error) {
	if ticket == nil {
		return nil, ErrorNil
	}
	if len(a.area) != 0 {
		for _, lot := range a.area {
			if _, exist := lot.ParkingField[ticket]; exist {
				return lot, nil
			}
		}
	}
	return nil, ErrorLocation
}

func (a *Attendant) UnPark(ticket *entity.Ticket) (*entity.Car, error) {
	if len(a.area) != 0 {
		for _, lot := range a.area {
			car, err := lot.Unpark(ticket)
			if err != nil {
				return nil, err
			}
			return car, nil
		}
	}
	return nil, ErrorNilLot
}
func NewAttendant(lots []*ParkingLot) *Attendant {
	attendant := &Attendant{}
	for _, lot := range lots {
		attendant.area = append(attendant.area, lot)
		attendant.availablelot = append(attendant.availablelot, lot)
		lot.Register(attendant)
	}
	attendant.style = StyleDefault()
	return attendant
}
func deleteLot(lots []*ParkingLot, lot *ParkingLot) []*ParkingLot {
	var result []*ParkingLot
	for _, value := range lots {
		if value != lot {
			result = append(result, value)
		}
	}
	return result
}
func (a *Attendant) Area() (lots []*ParkingLot) {
	lots = a.area
	return
}
