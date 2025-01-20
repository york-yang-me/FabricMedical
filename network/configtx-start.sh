#!/bin/bash
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=${PWD}/hyperledger-fabric-linux-amd64-2.4.2/bin:$PATH
fi

echo "1、clean the environment"
./stop.sh

echo "2、generate certificate and key(MSP materials), the generate result will be saved in the [crypto-config] file"
cryptogen generate --config=./crypto-config.yaml

echo "3、create orderer channel/genesis block"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel

echo "4、generate channel configuration tx-'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "5、define anchor peer for hospital"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/HospitalAnchor.tx -channelID appchannel -asOrg Hospital

echo "6、define anchor peer for patient"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/PatientAnchor.tx -channelID appchannel -asOrg Patient