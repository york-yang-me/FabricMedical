package verification

import (
	"bytes"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type Circuit struct {
	// struct tag on a variable is optional
	// default uses variable name and secret visibility.
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

// Define declares the circuit's constraints
// Hash = mimc(PreImage)
func (circuit *Circuit) Define(api frontend.API) error {
	// hash function
	mc, _ := mimc.NewMiMC(api)

	// specify constraints
	// mimc(preImage) == hash
	mc.Write(circuit.PreImage)
	api.AssertIsEqual(circuit.Hash, mc.Sum())

	return nil
}

// VerifyProof verify the proof
func VerifyProof(hash string, verifyKey, proofBytes []byte) (bool, error) {
	assignment := Circuit{Hash: hash}

	publicWitness1, err := frontend.NewWitness(&assignment, ecc.BLS12_381, frontend.PublicOnly())
	if err != nil {
		return false, err
	}

	proof := groth16.NewProof(ecc.BLS12_381)
	if _, err := proof.ReadFrom(bytes.NewBuffer(proofBytes)); err != nil {
		return false, err
	}

	vk := groth16.NewVerifyingKey(ecc.BLS12_381)
	if _, err := vk.ReadFrom(bytes.NewBuffer(verifyKey)); err != nil {
		return false, err
	}

	err = groth16.Verify(proof, vk, publicWitness1)
	if err != nil {
		return false, err
	}
	return true, nil
}
