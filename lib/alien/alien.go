package alien

import (
  "math/rand"
  "time"
)

type Alien struct {
  Name string
  Alive bool
  Trapped bool
  Steps int
  Curcity string
}

func (a *Alien) Destinations() []string {
  array := []string{"N", "S", "E", "W"}
	return random(array)
}

func random(array []string) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func (a *Alien) Die() {
  a.Alive = false
}

func (a *Alien) MoveTo(newCity string) {
  a.Curcity = newCity
  a.Steps++
}
