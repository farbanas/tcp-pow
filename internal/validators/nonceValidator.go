package validators

import "tcp-pow/internal/models"

// NonceValidator checks if nonce is valid. This is needed to prevent clients from sending whichever nonce they want
// and a solution for that nonce. We store the address and his nonce to the cache to prevent fraud.
type NonceValidator struct {
	cache Cache
}

func NewNonceValidator(cache Cache) *NonceValidator {
	return &NonceValidator{cache: cache}
}

// IsNonceValid checks if the nonce is correct for the given address. It does so by checking what's saved in the cache.
// If there's nothing saved in the cache for that address we assume that's ok and that it happened because this is the
// initial request, while a nonce has not been assigned to the client yet.
func (n *NonceValidator) IsNonceValid(tcpData models.TCPData) bool {
	nonce, err := n.cache.Get(tcpData.Address)
	if err != nil {
		// nonce is valid since there is no nonce for that address yet
		return true
	}

	if nonce != tcpData.Nonce {
		return false
	}

	return true
}
