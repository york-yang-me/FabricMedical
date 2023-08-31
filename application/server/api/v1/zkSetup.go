package v1

import (
	"application/pkg/circuit"
	"bytes"
	"encoding/hex"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log"
)

func Generate(preImage string, hash string) (string, string, error) {
	var c circuit.Circuit
	ccs, err := frontend.Compile(ecc.BLS12_381, r1cs.NewBuilder, &c)
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

	witness, _ := frontend.NewWitness(&c, ecc.BLS12_381)

	proof, _ := groth16.Prove(ccs, pk, witness)

	var proofBuffer bytes.Buffer
	proofBuffer.Reset()
	proof.WriteRawTo(&proofBuffer)
	proofBufferString := hex.EncodeToString(proofBuffer.Bytes())

	var vkBuffer bytes.Buffer
	vkBuffer.Reset()
	vk.WriteRawTo(&vkBuffer)
	vkBufferString := hex.EncodeToString(vkBuffer.Bytes())

	return vkBufferString, proofBufferString, nil
}
