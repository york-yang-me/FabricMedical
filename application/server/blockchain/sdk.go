package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// settings
var (
	sdk           *fabsdk.FabricSDK                                     // Fabric SDK
	configPath    = "config.yaml"                                       // Set file path
	channelName   = "appchannel"                                        // Channel name
	user          = "Admin"                                             // users
	chainCodeName = "fabric-medical"                                    // chaincode name
	endpoints     = []string{"peer0.patient.com", "peer0.hospital.com"} // the nodes need to send transaction
)

// Init
func Init() {
	var err error
	// use setting files to init SDK
	sdk, err = fabsdk.New(config.FromFile(configPath))
	if err != nil {
		panic(err)
	}
}

// ChannelExecute BC interactive
func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
	// create client, show the identity in channel
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// execute writing operation to BC Ledger （use chaincode invoke）
	resp, err := cli.Execute(channel.Request{
		ChaincodeID: chainCodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(endpoints...))
	if err != nil {
		return channel.Response{}, err
	}
	// return the result after executed
	return resp, nil
}

// ChannelQuery BC query
func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
	// create client, show the identity in channel
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// execute querying operation to BC Ledger (use chaincode invoke), only return the result
	resp, err := cli.Query(channel.Request{
		ChaincodeID: chainCodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(endpoints...))
	if err != nil {
		return channel.Response{}, err
	}
	// return the result after executed
	return resp, nil
}
