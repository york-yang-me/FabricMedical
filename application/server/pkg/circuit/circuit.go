package circuit

import (
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
