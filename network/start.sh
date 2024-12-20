#!/bin/bash
echo "Blockchain: Link Start!"
docker-compose -f docker-compose.yaml up -d
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
HospitalPeer2Cli="CORE_PEER_ADDRESS=peer2.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer2.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer2.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer2.hospital.com/tls/ca.crt"
HospitalPeer3Cli="CORE_PEER_ADDRESS=peer3.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer3.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer3.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer3.hospital.com/tls/ca.crt"
HospitalPeer4Cli="CORE_PEER_ADDRESS=peer4.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/ca.crt"
HospitalPeer5Cli="CORE_PEER_ADDRESS=peer5.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/ca.crt"
PatientPeer0Cli="CORE_PEER_ADDRESS=peer0.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer0.patient.com/tls/ca.crt"
PatientPeer1Cli="CORE_PEER_ADDRESS=peer1.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer1.patient.com/tls/ca.crt"
PatientPeer2Cli="CORE_PEER_ADDRESS=peer2.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer2.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer2.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer2.patient.com/tls/ca.crt"
PatientPeer3Cli="CORE_PEER_ADDRESS=peer3.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer3.patient.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer3.patient.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer3.patient.com/tls/ca.crt"
OrdererCa="/etc/hyperledger/orderer/gmp.com/tlsca/tlsca.gmp.com-cert.pem"

echo "7、create channel"
docker exec cli bash -c "$HospitalPeer0Cli peer channel create -o orderer.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/appchannel.tx --cafile $OrdererCa"

echo "8、add all notes to channel"
docker exec cli bash -c "$HospitalPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer2Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer3Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer4Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer5Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$PatientPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$PatientPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$PatientPeer2Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$PatientPeer3Cli peer channel join -b appchannel.block"


echo "9、Update anchor notes"
docker exec cli bash -c "$HospitalPeer0Cli peer channel update -o orderer.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/HospitalAnchor.tx  --cafile $OrdererCa"
docker exec cli bash -c "$PatientPeer0Cli peer channel update -o orderer.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/PatientAnchor.tx --cafile $OrdererCa"