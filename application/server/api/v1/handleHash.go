package v1

import (
	"github.com/consensys/gnark-crypto/hash"
	"math/big"
)

// HashCalc  calculate mimc hash by BLS12_381
func HashCalc(preImage string) string {
	imageHash, _ := big.NewInt(0).SetString(preImage, 10)

	h := hash.MIMC_BLS12_381.New()
	h.Write(imageHash.Bytes())
	rd := h.Sum(nil)
	result := big.NewInt(0).SetBytes(rd).String()

	return result
}
