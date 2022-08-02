package factories

import (
	"tcp-pow/internal/validators"
	"tcp-pow/internal/validators/tests/mocks"

	"github.com/golang/mock/gomock"
)

func CreateNonceValidator(ctrl *gomock.Controller) (*mocks.MockCache, *validators.NonceValidator) {
	cache := mocks.NewMockCache(ctrl)
	nonceValidator := validators.NewNonceValidator(cache)

	return cache, nonceValidator
}
