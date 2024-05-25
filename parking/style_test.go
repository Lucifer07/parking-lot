package parking_test

import (
	"testing"

	"github.com/Lucifer07/parking-lot/entity"
	"github.com/Lucifer07/parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStyle(t *testing.T) {
	t.Run("park using Default style successed", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(1)
		car := entity.AddCar("hhhh")
		parkingArea := []*parking.ParkingLot{lot, lot2}
		attendant := parking.NewAttendant(parkingArea)
		attendant.ChangeStyle(parking.StyleDefault())
		ticket, _ := attendant.Park(car)
		car2, _ := attendant.UnPark(ticket)
		assert.Equal(t, car, car2)
	})

	t.Run("park using Default style should be give error if all lot is full", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(1)
		car := entity.AddCar("hhhh")
		car2 := entity.AddCar("hhhh")
		car3 := entity.AddCar("hhhh")
		parkingArea := []*parking.ParkingLot{lot, lot2}
		attendant := parking.NewAttendant(parkingArea)
		attendant.Park(car)
		attendant.Park(car2)
		_, err := attendant.Park(car3)
		assert.EqualError(t, err, parking.ErrorParkingFull.Error())
	})
}

func TestCapacityStyle(t *testing.T) {
	t.Run("park using capacity style successed and check location", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(2)
		car := entity.AddCar("hhhh")
		parkingArea := []*parking.ParkingLot{lot, lot2}
		attendant := parking.NewAttendant(parkingArea)
		attendant.ChangeStyle(parking.StyleCapacity())
		ticket, _ := attendant.Park(car)
		lot3, _ := attendant.Location(ticket)
		assert.Equal(t, lot2, lot3)
	})
	t.Run("park using capacity style should be give error if all lot is full", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(1)
		parkingArea := []*parking.ParkingLot{lot, lot2}
		car1 := entity.AddCar("hhhh")
		car2 := entity.AddCar("hhhh")
		car3 := entity.AddCar("hhhh")
		attendant := parking.NewAttendant(parkingArea)
		attendant.ChangeStyle(parking.StyleCapacity())
		attendant.Park(car1)
		attendant.Park(car2)
		_, err := attendant.Park(car3)
		assert.EqualError(t, err, parking.ErrorParkingFull.Error())
	})
}
func TestFreeSpaceStyle(t *testing.T) {
	t.Run("park using FreeSpace style successed and check location", func(t *testing.T) {
		lot := parking.NewParkingLot(2)
		lot2 := parking.NewParkingLot(3)
		car := entity.AddCar("hhhh")
		car1 := entity.AddCar("hhhh")
		car2 := entity.AddCar("hhhh")
		parkingArea := []*parking.ParkingLot{lot, lot2}
		attendant := parking.NewAttendant(parkingArea)
		attendant.ChangeStyle(parking.StyleFreeSpace())
		ticket, _ := attendant.Park(car)
		attendant.Park(car1)
		attendant.UnPark(ticket)
		ticket2, _ := attendant.Park(car2)
		lot3, _ := attendant.Location(ticket2)
		assert.Equal(t, lot2, lot3)
	})
	t.Run("park using FreeSpace style should be give error if all lot is full", func(t *testing.T) {
		lot := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(1)
		car := entity.AddCar("hhhh")
		car2 := entity.AddCar("hhhh")
		car3 := entity.AddCar("hhhh")
		parkingArea := []*parking.ParkingLot{lot, lot2}
		attendant := parking.NewAttendant(parkingArea)
		attendant.ChangeStyle(parking.StyleFreeSpace())
		attendant.Park(car)
		attendant.Park(car2)
		_, err := attendant.Park(car3)
		assert.EqualError(t, err, parking.ErrorParkingFull.Error())
	})
}
