package verification

import (
	"bytes"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

// VerifyProof verify the proof
func VerifyProof(hash string, verifyKey, proofBytes []byte) (bool, error) {
	assignment1 := Circuit{
		Hash: hash,
	}
	publicWitness1, err := frontend.NewWitness(&assignment1, ecc.BN254, frontend.PublicOnly())
	if err != nil {
		return false, err
	}

	proof := groth16.NewProof(ecc.BN254)
	if _, err := proof.ReadFrom(bytes.NewBuffer(proofBytes)); err != nil {
		return false, err
	}

	vk := groth16.NewVerifyingKey(ecc.BN254)
	if _, err := vk.ReadFrom(bytes.NewBuffer(verifyKey)); err != nil {
		return false, err
	}

	err = groth16.Verify(proof, vk, publicWitness1)
	if err != nil {
		return false, err
	}
	return true, nil
}
