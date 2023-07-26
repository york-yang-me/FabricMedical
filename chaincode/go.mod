module chaincode

go 1.20

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20201119163726-f8ef75b17719
	github.com/hyperledger/fabric-protos-go v0.0.0-20211118165945-23d738fc3553
)

require (
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/stretchr/testify v1.7.1-0.20210116013205-6990a05d54c2 // indirect; includes ErrorContains
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210420205809-ac73e9fd8988 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/onsi/gomega => github.com/onsi/gomega v1.9.0

replace github.com/cespare/xxhash/v2 => github.com/cespare/xxhash/v2 v2.1.2

// fix for Go 1.17 in github.com/prometheus/client_golang dependency without updating protobuf
