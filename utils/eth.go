package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func GenerateAddressFromDenom(denom string) (common.Address, error) {
	hash := sha3.NewLegacyKeccak256()
	if _, err := hash.Write([]byte(denom)); err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(hash.Sum(nil)), nil
}
