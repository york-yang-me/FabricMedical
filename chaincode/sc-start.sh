#!/bin/bash

HospitalPeer0Cli="CORE_PEER_ADDRESS=peer0.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/ca.crt"
InstitutePeer0Cli="CORE_PEER_ADDRESS=peer0.institute.com:7051 CORE_PEER_LOCALMSPID=InstituteMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/institute.com/users/Admin@institute.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/ca.crt"
OrdererCa="/etc/hyperledger/orderer/government.com/tlsca/tlsca.government.com-cert.pem"


# Package chaincode
echo "10、package the chaincode"
docker exec cli bash -c "peer lifecycle chaincode package medical_chaincode.tar.gz --path src/chaincode -l golang --label medical_chaincode"

# -n chaincode name
# -v version number
# -p chaincode category
echo "11、install chaincode"
docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode install medical_chaincode.tar.gz"
docker exec cli bash -c "$InstitutePeer0Cli peer lifecycle chaincode install medical_chaincode.tar.gz"

echo "12. check the chaincode installed"
docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode queryinstalled"
# shellcheck disable=SC2046
PACKAGE_ID=$( echo `docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode queryinstalled"`| awk 'split($7, arr, ",") {print arr[1]}')

# Deploy chaincode
# -n the name of installed chaincode
# -v version number
# -C channel  one channel == one chain
# -c parameter incoming init parameter
echo "13、deploy chaincode"
docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode approveformyorg -o orderer.government.com:7050 --signature-policy \"OR ('HospitalMSP.member','InstituteMSP.member')\" --tls --cafile $OrdererCa  --channelID appchannel --name fabric-medical --version 1.0.0 --package-id $PACKAGE_ID --sequence 1 --waitForEvent"
docker exec cli bash -c "$InstitutePeer0Cli peer lifecycle chaincode approveformyorg -o orderer.government.com:7050 --signature-policy \"OR ('HospitalMSP.member','InstituteMSP.member')\" --tls --cafile $OrdererCa  --channelID appchannel --name fabric-medical --version 1.0.0 --package-id $PACKAGE_ID --sequence 1 --waitForEvent"
docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode checkcommitreadiness --channelID appchannel --name fabric-medical --version 1.0 --sequence 1 --output json"
docker exec cli bash -c "$HospitalPeer0Cli peer lifecycle chaincode commit -o orderer.government.com:7050 --signature-policy \"OR ('HospitalMSP.member','InstituteMSP.member')\" --tls --cafile $OrdererCa --channelID appchannel --name fabric-medical --version 1.0.0 --sequence 1 --peerAddresses peer0.hospital.com:7051 --tlsRootCertFiles /etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/ca.crt --peerAddresses peer0.institute.com:7051 --tlsRootCertFiles /etc/hyperledger/peer/institute.com/peers/peer0.institute.com/tls/ca.crt"

echo "waiting for instantiating chaincode, countdown 5 seconds"
sleep 5

# interact with chaincode, verify chaincode if it's correctly installed and check the blockchain network
echo "14、verify the chaincode"
docker exec cli bash -c "$HospitalPeer0Cli peer chaincode invoke -C appchannel -n fabric-medical -c '{\"Args\":[\"hello\"]}' --tls --cafile $OrdererCa"
docker exec cli bash -c "$InstitutePeer0Cli peer chaincode invoke -C appchannel -n fabric-medical -c '{\"Args\":[\"hello\"]}' --tls --cafile $OrdererCa"
