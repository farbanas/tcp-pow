package cache

import "fmt"

// Cache is a simple in-memory map with Get and Set methods. Operates only with strings.
// It is meant to be easily replacable with a real implementation like Redis.
type Cache struct {
	localData map[string]string
}

func NewCache() *Cache {
	return &Cache{
		localData: make(map[string]string),
	}
}

// Get gets the value saved under the given address. If it doesn't exist an error is returned.
func (c *Cache) Get(address string) (string, error) {
	data, ok := c.localData[address]
	if !ok {
		return "", fmt.Errorf("data for address %s missing", address)
	}

	return data, nil
}

// Set sets the nonce under the given address in the map.
func (c *Cache) Set(address string, nonce string) {
	c.localData[address] = nonce
}
