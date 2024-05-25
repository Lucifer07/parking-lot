package entity_test

import (
	"testing"

	. "github.com/Lucifer07/parking-lot/entity"
	"github.com/stretchr/testify/assert"
)

func TestAddCar(t *testing.T) {
	t.Run("should return car", func(t *testing.T) {
		t1 := AddCar("xxxxx")
		assert.NotEmpty(t, t1)
	})
}
