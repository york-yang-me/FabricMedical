module chaincode

go 1.20

require (
	github.com/consensys/gnark v0.7.0
	github.com/consensys/gnark-crypto v0.7.0
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20201119163726-f8ef75b17719
	github.com/hyperledger/fabric-protos-go v0.0.0-20211118165945-23d738fc3553
)

require (
	github.com/fxamacker/cbor/v2 v2.2.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/rs/zerolog v1.26.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
)

require (
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.31.0 // indirect
)

replace github.com/onsi/gomega => github.com/onsi/gomega v1.9.0

replace github.com/cespare/xxhash/v2 => github.com/cespare/xxhash/v2 v2.1.2

// fix for Go 1.17 in github.com/prometheus/client_golang dependency without updating protobuf
