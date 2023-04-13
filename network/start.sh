#!/bin/bash
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=/usr/local/bin:$PATH
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

echo "6、define anchor peer for institute"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/InstituteAnchor.tx -channelID appchannel -asOrg Institute

echo "Blockchain: Link Start!"
docker-compose up -d
echo "waiting for nodes start-up to complete, countdown 10 seconds"
sleep 10

# Open two-way authentication with TLS
HospitalPeer0Cli="CORE_PEER_ADDRESS=peer0.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/ca.crt"
HospitalPeer1Cli="CORE_PEER_ADDRESS=peer1.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer1.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer1.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer1.hospital.com/tls/ca.crt"
InstitutePeer0Cli="CORE_PEER_ADDRESS=peer0.institute.com:7051 CORE_PEER_LOCALMSPID=InstituteMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/institute.com/users/Admin@institute.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/ca.crt"
InstitutePeer1Cli="CORE_PEER_ADDRESS=peer1.institute.com:7051 CORE_PEER_LOCALMSPID=InstituteMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/institute.com/users/Admin@institute.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer1.institute.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/institute.com/peers/peer1.institute.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer1.institute.com/tls/ca.crt"
OrdererCa="/etc/hyperledger/orderer/government.com/tlsca/tlsca.government.com-cert.pem"

echo "7、create channel"
docker exec cli bash -c "$HospitalPeer0Cli peer channel create -o orderer.government.com:7050 --tls -c appchannel -f /etc/hyperledger/config/appchannel.tx --cafile $OrdererCa"

echo "8、add all notes to channel"
docker exec cli bash -c "$HospitalPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$InstitutePeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$InstitutePeer1Cli peer channel join -b appchannel.block"

echo "9、Update anchor notes"
docker exec cli bash -c "$HospitalPeer0Cli peer channel update -o orderer.government.com:7050 --tls -c appchannel -f /etc/hyperledger/config/HospitalAnchor.tx  --cafile $OrdererCa"
docker exec cli bash -c "$InstitutePeer0Cli peer channel update -o orderer.government.com:7050 --tls -c appchannel -f /etc/hyperledger/config/InstituteAnchor.tx --cafile $OrdererCa"