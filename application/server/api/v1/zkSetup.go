package v1

import (
	"application/pkg/circuit"
	"bytes"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log"
)

func Generate(preImage string, hash string) ([]byte, []byte, error) {
	var c circuit.Circuit
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &c)
	if err != nil {
		log.Fatalln("generate circuit failed:", err)
	}

	// groth16 zkSNARK: SetUp
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatalln("groth16 setup failed")
	}

	c.PreImage = preImage
	c.Hash = hash

	witness, err := frontend.NewWitness(&c, ecc.BN254)
	if err != nil {
		return nil, nil, fmt.Errorf("generate witness error: %s", err)
	}

	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		return nil, nil, fmt.Errorf("generate proof error: %s", err)
	}

	var proofBuffer bytes.Buffer
	proofBuffer.Reset()
	proof.WriteRawTo(&proofBuffer)
	proofBufferBytes := proofBuffer.Bytes()

	var vkBuffer bytes.Buffer
	vkBuffer.Reset()
	vk.WriteRawTo(&vkBuffer)
	vkBufferBytes := vkBuffer.Bytes()

	return vkBufferBytes, proofBufferBytes, nil
}
