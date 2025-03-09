package luckshot

import (
	"math/rand"
)

type LuckShot struct {
	shot bool
}

func (ls *LuckShot) Do(fn func()) {
	if ls.shot {
		fn()
	}
}

func Chance(base, target int) *LuckShot {
	randNum := rand.Intn(base)
	shot := randNum < target
	return &LuckShot{shot}
}
