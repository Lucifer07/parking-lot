package parking_test

import (
	"testing"

	"github.com/Lucifer07/parking-lot/entity"
	"github.com/Lucifer07/parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestNewAttendant(t *testing.T) {
	t.Run("should be return same type if successed", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := &parking.Attendant{}
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		assert.IsType(t, lot2, attendant)
	})
}
func TestAddParkinglot(t *testing.T) {
	t.Run("should be return value if successed", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		attendant = attendant.AddParkinglot(1)
		totalLots := len(attendant.Area())
		assert.Equal(t, 2, totalLots)
	})

}
func TestParkAttendant(t *testing.T) {
	t.Run("should return parkingfull cause dont have space", func(t *testing.T) {
		car := entity.AddCar("aaaa")
		target := parking.ErrorParkingFull
		parkingArea := []*parking.ParkingLot{}
		attendant := parking.NewAttendant(parkingArea)
		_, err := attendant.Park(car)
		assert.Equal(t, target, err)
	})
	t.Run("should be give error if all lot is full", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		car := entity.AddCar("hhhh")
		car2 := entity.AddCar("bbb")
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		attendant.Park(car)
		_, err := attendant.Park(car2)
		assert.Equal(t, parking.ErrorParkingFull, err)
	})
	t.Run("should be return same pointer", func(t *testing.T) {
		lot := parking.NewParkingLot(2)
		car := entity.AddCar("hhhh")
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		ticket, _ := attendant.Park(car)
		car2, _ := attendant.UnPark(ticket)
		assert.Equal(t, car, car2)
	})
}
func TestUnParkAttendant(t *testing.T) {
	t.Run("should be error if dont have area", func(t *testing.T) {
		parkingArea := []*parking.ParkingLot{}
		ticket := entity.NewTicket()
		attendant := parking.NewAttendant(parkingArea)
		_, err := attendant.UnPark(&ticket)
		assert.Equal(t, parking.ErrorNilLot, err)
	})
	t.Run("should be error if using ticket was used", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		car := entity.AddCar("hhhh")
		ticket, _ := attendant.Park(car)
		attendant.UnPark(ticket)
		_, err := attendant.UnPark(ticket)
		assert.Equal(t, parking.ErrorUnrecognized, err)
	})
	t.Run("should be error if using wrong ticket", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		car := entity.AddCar("hhhh")
		attendant.Park(car)
		ticket := entity.NewTicket()
		_, err := attendant.UnPark(&ticket)
		assert.Equal(t, parking.ErrorUnrecognized, err)
	})

	t.Run("should no error if successed", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		attendant := parking.NewAttendant([]*parking.ParkingLot{lot})
		car := entity.AddCar("hhhh")
		ticket, _ := attendant.Park(car)
		car2, _ := attendant.UnPark(ticket)
		assert.Equal(t, car, car2)
	})
}
