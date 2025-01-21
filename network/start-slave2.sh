#!/bin/bash
echo "Blockchain: Link Start!"
docker-compose -f docker-compose-slave2.yaml up -d
echo "waiting for nodes start-up to complete, countdown 10 seconds"
sleep 10

# Open two-way authentication with TLS
HospitalPeer4Cli="CORE_PEER_ADDRESS=peer4.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer4.hospital.com/tls/ca.crt"
HospitalPeer5Cli="CORE_PEER_ADDRESS=peer5.hospital.com:7051 CORE_PEER_LOCALMSPID=HospitalMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/hospital.com/users/Admin@hospital.com/msp \
                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/server.crt \
                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/server.key \
                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/hospital.com/peers/peer5.hospital.com/tls/ca.crt"
#PatientPeer4Cli="CORE_PEER_ADDRESS=peer4.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
#                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer4.patient.com/tls/server.crt \
#                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer4.patient.com/tls/server.key \
#                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer4.patient.com/tls/ca.crt"
#PatientPeer5Cli="CORE_PEER_ADDRESS=peer5.patient.com:7051 CORE_PEER_LOCALMSPID=PatientMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/patient.com/users/Admin@patient.com/msp \
#                  CORE_PEER_TLS_ENABLED=true CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer5.patient.com/tls/server.crt \
#                  CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/patient.com/peers/peer5.patient.com/tls/server.key \
#                  CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/patient.com/peers/peer5.patient.com/tls/ca.crt"
OrdererCa="/etc/hyperledger/orderer/gmp.com/tlsca/tlsca.gmp.com-cert.pem"

echo "7、create channel"
docker exec cli bash -c "$HospitalPeer4Cli peer channel create -o orderer3.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/appchannel.tx --cafile $OrdererCa"

echo "8、add all notes to channel"
docker exec cli bash -c "$HospitalPeer4Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$HospitalPeer5Cli peer channel join -b appchannel.block"

echo "9、Update anchor notes"
docker exec cli bash -c "$HospitalPeer4Cli peer channel update -o orderer3.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/HospitalAnchor.tx  --cafile $OrdererCa"
#docker exec cli bash -c "$PatientPeer4Cli peer channel update -o orderer3.gmp.com:7050 --tls -c appchannel -f /etc/hyperledger/config/PatientAnchor.tx --cafile $OrdererCa"