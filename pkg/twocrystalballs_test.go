package day1

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	t.Run("first_test", func(t *testing.T) {
		idx := rand.Intn(10000)
		data := make([]bool, 10000)

		for i := idx; i < 10000; i++ {
			data[i] = true
		}

		val := TwoCrystalBalls(data)
		if val != idx {
			t.Errorf("expected %d | got %d\n", idx, val)
		}
	})

	t.Run("second_test", func(t *testing.T) {
		data := make([]bool, 821)

		val := TwoCrystalBalls(data)

		if val != -1 {
			t.Errorf("expected: %d | got: %d\n", -1, val)
		}
	})
}
