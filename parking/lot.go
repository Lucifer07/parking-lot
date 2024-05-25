package parking

import (
	"errors"

	"github.com/Lucifer07/parking-lot/entity"
)

var (
	ErrorNil          = errors.New("nil value")
	ErrorParkedTwice  = errors.New("should not able parked twice")
	ErrorUnrecognized = errors.New("unrecognized parking ticket")
)

const (
	Full      = "full"
	Available = "available"
)

type Publisher interface {
	Register(Observer)
	NotifyObserver(condition string)
}

func (p *ParkingLot) Register(o Observer) {
	p.subscribers = append(p.subscribers, o)
}
func (p *ParkingLot) NotifyObserver(condition string) {
	for _, v := range p.subscribers {
		if condition != Full {
			v.UpdateAvailable(p)
			return
		}
		v.UpdateFull(p)
	}
}

type ParkingLot struct {
	max          int
	ParkingField map[*entity.Ticket]*entity.Car
	subscribers  []Observer
}

func NewParkingLot(max int) *ParkingLot {
	return &ParkingLot{
		max:          max,
		ParkingField: make(map[*entity.Ticket]*entity.Car),
	}
}
func (p *ParkingLot) Park(car *entity.Car) (*entity.Ticket, error) {
	if car != nil && p != nil {
		if p.max > len(p.ParkingField) {
			if p.max >= len(p.ParkingField) {
				ticket := entity.NewTicket()
				p.ParkingField[&ticket] = car
				if p.max == len(p.ParkingField) {
					p.NotifyObserver(Full)
				}
				return &ticket, nil
			}
		}
		return nil, ErrorParkingFull
	}
	return nil, ErrorNil

}

func (p *ParkingLot) Unpark(ticket *entity.Ticket) (*entity.Car, error) {
	if ticket != nil {
		if _, exist := p.ParkingField[ticket]; !exist {
			return nil, ErrorUnrecognized
		}
		p.NotifyObserver(Available)
		car := p.ParkingField[ticket]
		delete(p.ParkingField, ticket)
		return car, nil
	}
	return nil, ErrorNil
}
func IsParked(car *entity.Car, lot *ParkingLot) bool {
	if car != nil && lot != nil {
		for _, parkedCar := range lot.ParkingField {
			if parkedCar == car {
				return true
			}
		}
	}
	return false
}
