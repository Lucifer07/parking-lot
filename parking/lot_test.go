package parking_test

import (
	"testing"

	"github.com/Lucifer07/parking-lot/entity"
	"github.com/Lucifer07/parking-lot/parking"
	"github.com/Lucifer07/parking-lot/parking/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewParkingLot(t *testing.T) {
	t.Run("should return ticket with non empty ID", func(t *testing.T) {
		lot := parking.NewParkingLot(10)
		assert.NotEmpty(t, lot)
	})
}
func TestPark(t *testing.T) {

	t.Run("should be return error nil parameter equal nil ", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		target := parking.ErrorNil
		_, err := lot.Park(nil)
		assert.Equal(t, target, err)
	})
	t.Run("should be return error parkingfull if out of capacity ", func(t *testing.T) {
		target := parking.ErrorParkingFull
		lot := parking.NewParkingLot(1)
		c1 := entity.AddCar("aaa")
		c2 := entity.AddCar("bbb")
		_, _ = lot.Park(c1)
		_, err := lot.Park(c2)
		assert.Equal(t, target, err)
	})

	t.Run("should give return same pointer with car", func(t *testing.T) {
		lot := parking.NewParkingLot(2)
		car := entity.AddCar("aaa")
		ticket, _ := lot.Park(car)
		car2, _ := lot.Unpark(ticket)
		assert.Equal(t, car2, car)
	})
}
func TestUnpark(t *testing.T) {
	t.Run("should be return error nil", func(t *testing.T) {
		target := parking.ErrorNil
		lot := parking.NewParkingLot(1)
		_, err := lot.Unpark(nil)
		assert.Equal(t, target, err)
	})

	t.Run("should be return error ErrorUnrecognized if using wrong ticket ", func(t *testing.T) {
		target := parking.ErrorUnrecognized
		lot := parking.NewParkingLot(1)
		car := entity.AddCar("aaaa")
		ticket := entity.NewTicket()
		lot.Park(car)
		_, err := lot.Unpark(&ticket)
		assert.Equal(t, target, err)
	})

	t.Run("should be return error ErrorUnrecognized if using ticket was used ", func(t *testing.T) {
		target := "unrecognized parking ticket"
		lot := parking.NewParkingLot(1)
		car := entity.AddCar("aaaa")
		ticket, _ := lot.Park(car)
		lot.Unpark(ticket)
		_, err := lot.Unpark(ticket)
		assert.Equal(t, target, err.Error())
	})
	t.Run("should return same pointer with car", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		car := entity.AddCar("aaaa")
		ticket, _ := lot.Park(car)
		car1, _ := lot.Unpark(ticket)
		assert.Equal(t, car, car1)
	})

}
func TestNotifyFull(t *testing.T) {
	lot := parking.NewParkingLot(1)
	MockO := new(mocks.Observer)
	MockO.On("UpdateFull", lot).Return()
	lot.Register(MockO)
	lot.NotifyObserver(parking.Full)
	lot.NotifyObserver(parking.Full)
	MockO.AssertNumberOfCalls(t, "UpdateFull", 2)
}
func TestNotifyAvailable(t *testing.T) {
	lot := parking.NewParkingLot(1)
	MockO := new(mocks.Observer)
	MockO.On("UpdateAvailable", lot).Return()
	lot.Register(MockO)
	lot.NotifyObserver(parking.Available)
	lot.NotifyObserver(parking.Available)
	MockO.AssertNumberOfCalls(t, "UpdateAvailable", 2)
}
