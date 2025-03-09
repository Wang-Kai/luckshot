package luckshot

import "testing"

func TestLuckshot(t *testing.T) {
	Chance(10, 1).Do(func() {
		println("There is a 1/10 probability of printing it out.")
	})
}
