//Package token_generator provides functionality to generate unique random strings.
package token_generator

import (
	"math/rand"
	"time"
)

/*
Main Generator structure
	n		Number of generated strings to be cached
	cache_length	Number of characters in the string
	cache		Array of generated strings
*/
type Generator struct {
	n            int
	cache_length int
	cache        []string
}

/*
Initializer for the token generator. To initialize with custom settings set:-
	n		Any positive integer value (default 1000)
	cache_length	Any positive integer value (default 200)
*/
func (g *Generator) New() {
	if g.n == 0 {
		g.n = 1000
	}
	if g.cache_length == 0 {
		g.cache_length = 200
	}
	g.cache = make([]string, g.n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < g.n; i++ {
		g.cache[i] = RandomString(g.cache_length)
	}
}

// Token Generator function
func (g Generator) GetToken(size int) string {
	i := rand.Intn(g.n)
	start := rand.Intn(g.cache_length - size - 1)
	return g.cache[i][start : start+size]
}
