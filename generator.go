//Package token_generator provides functionality to generate unique random strings.
package token_generator

import (
	"math/rand"
	"time"
)

/*
Main Generator structure
	N		Number of generated strings to be cached
	Cache_length	Number of characters in the string
	cache		Array of generated strings
*/
type Generator struct {
	N            int
	Cache_length int
	cache        []string
}

/*
Initializer for the token generator. To initialize with custom settings set:-
	N		Any positive integer value (default 1000)
	Cache_length	Any positive integer value (default 200)
*/
func (g *Generator) New() {
	if g.N == 0 {
		g.N = 1000
	}
	if g.Cache_length == 0 {
		g.Cache_length = 200
	}
	g.cache = make([]string, g.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < g.N; i++ {
		g.cache[i] = RandomString(g.Cache_length)
	}
}

// Token Generator function
func (g Generator) GetToken(size int) string {
	i := rand.Intn(g.N)
	start := rand.Intn(g.Cache_length - size - 1)
	return g.cache[i][start : start+size]
}
