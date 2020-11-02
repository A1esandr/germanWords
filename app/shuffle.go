package app

import (
	"math/rand"
)

func Shuffle(w [][]string) {
	for i := range w {
		j := rand.Intn(i + 1)
		w[i], w[j] = w[j], w[i]
	}
}
