package tokengenerator

//Package token_generator provides functionality to generate unique random strings.

import (
	"math/rand"
	"time"
)

/*
Main Generator structure
	N		Number of generated strings to be cached
	CacheLength	Number of characters in the string
	cache		Array of generated strings
*/
type Generator struct {
	N           int
	CacheLength int
	cache       []string
}

/*
Initializer for the token generator. To initialize with custom settings set:-
	N		Any positive integer value (default 1000)
	CacheLength	Any positive integer value (default 200)
*/
func (g *Generator) New() {
	if g.N == 0 {
		g.N = 1000
	}
	if g.CacheLength == 0 {
		g.CacheLength = 200
	}
	g.cache = make([]string, g.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < g.N; i++ {
		g.cache[i] = RandomString(g.CacheLength)
	}
}

// Token Generator function
func (g Generator) GetToken(size int) string {
	i := rand.Intn(g.N)
	start := rand.Intn(g.CacheLength - size - 1)
	return g.cache[i][start : start+size]
}
