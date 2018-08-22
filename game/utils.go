package game

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var consonants = []rune("qwrtypsdfghjklzxcvbnm")
var vowels = []rune("aeiou")

// NameGenerator generates a random string alternating vowels and consonants
func NameGenerator() string {
	name := []string{}
	b := make([]rune, 3)
	c := make([]rune, 3)
	for i := range b {
		b[i] = consonants[rand.Intn(3)]
		c[i] = vowels[rand.Intn(3)]

	}
	name = append(name, string(b[0]), string(c[0]), string(b[1]), string(c[1]), string(b[2]), string(c[2]))
	return strings.Join(name, "")
}
