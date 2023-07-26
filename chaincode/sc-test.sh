#!/bin/bash

HospitalPeer0Cli="CORE_PEER_ADDRESS=peer0.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/ca.crt"
HospitalPeer1Cli="CORE_PEER_ADDRESS=peer0.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer0.hospital.com/tls/ca.crt"
PatientPeer0Cli="CORE_PEER_ADDRESS=peer0.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/ca.crt"
PatientPeer1Cli="CORE_PEER_ADDRESS=peer1.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/ca.crt"
OrdererCa="/etc/hyperledger/orderer/gmp.com/tlsca/tlsca.gmp.com-cert.pem"


#echo "15、Init the chaincode"
#docker exec cli bash -c "$HospitalPeer0Cli peer chaincode invoke -o orderer.gmp.com:7050 -C appchannel -n fabric-medical --isInit -c '{\"Args\":[\"init\"]}' --tls --cafile $OrdererCa"

# interact with chaincode, verify chaincode if it's correctly installed and check the blockchain network
echo "16、verify the chaincode"
docker exec cli bash -c "$HospitalPeer0Cli peer chaincode invoke -C appchannel -n fabric-medical -c '{\"Args\":[\"hello\"]}' --tls --cafile $OrdererCa"
docker exec cli bash -c "$PatientPeer0Cli peer chaincode invoke -C appchannel -n fabric-medical -c '{\"Args\":[\"hello\"]}' --tls --cafile $OrdererCa"