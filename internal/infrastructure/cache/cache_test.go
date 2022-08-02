package cache_test

import (
	"tcp-pow/internal/infrastructure/cache"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cache Suite")
}

var _ = Describe("Given a cache", func() {
	Describe("When get is called", func() {
		var c *cache.Cache
		var address string
		BeforeEach(func() {
			c = cache.NewCache()
			c.Set("a", "test1")
			c.Set("b", "test2")
			address = "a"
		})

		Context("And address doesn't exist in the cache", func() {
			BeforeEach(func() {
				address = "c"
			})

			It("Then cache should return an error", func() {
				_, err := c.Get(address)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("And address does exist in the cache", func() {
			It("Then cache should return the value stored under that address", func() {
				value, err := c.Get(address)
				Expect(err).To(BeNil())
				Expect(value).To(BeEquivalentTo("test1"))
			})
		})
	})

	Describe("When get is called", func() {
		var c *cache.Cache
		var address string
		var nonce string
		BeforeEach(func() {
			c = cache.NewCache()
			address = "a"
			nonce = "test1"
		})

		Context("And setting succeeds", func() {
			It("Then cache should contain the value that we've set", func() {
				c.Set(address, nonce)
				value, err := c.Get(address)
				Expect(err).To(BeNil())
				Expect(value).To(BeEquivalentTo(nonce))
			})
		})
	})
})
