package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"application/blockchain"
	"application/pkg/cron"
	"application/routers"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
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

func main() {
	timeLocal, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Printf("Failed to set time zone %s", err)
	}
	time.Local = timeLocal

	// define circuit
	var circuit Circuit
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Printf("Circuit compile failed: %s", err)
	}
	// groth16 zkSNARK: Setup
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Printf("Groth16 zkSNARK setup failed: %s", err)
	}
	log.Println(pk, vk)

	blockchain.Init()
	go cron.Init()

	endPoint := fmt.Sprintf("0.0.0.0:%d", 8000)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routers.InitRouter(),
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}

//// witness definition
//assignment :=  Circuit{
//	PreImage: "16130099170765464552823636852555369511329944820189892919423002775646948828469",
//	Hash:     "8674594860895598770446879254410848023850744751986836044725552747672873438975",
//}
//
//witness, err := frontend.NewWitness(&assignment, ecc.BN254)
//if err != nil {
//	panic(err)
//}
//
//proof, err := groth16.Prove(ccs, pk, witness)
//if err != nil {
//	panic(err)
//}
//var proofBuffer bytes.Buffer
//proofBuffer.Reset()
//proof.WriteRawTo(&proofBuffer)
//proofBuffer.Bytes()
//
//var vkBuffer bytes.Buffer
//vkBuffer.Reset()
//vk.WriteRawTo(&vkBuffer)
//vkBuffer.Bytes()
//
//fmt.Printf("proof: %s\n", hex.EncodeToString(proofBuffer.Bytes()))
//fmt.Printf("verifykey: %s\n", hex.EncodeToString(vkBuffer.Bytes()))
//
//// VerifyProof 函数放到智能合约执行
//r, err := VerifyProof("8674594860895598770446879254410848023850744751986836044725552747672873438975",
//	vkBuffer.Bytes(),
//	proofBuffer.Bytes())
//if err != nil {
//	panic(err)
//}
//
//fmt.Printf("verify result: %v\n", r)
