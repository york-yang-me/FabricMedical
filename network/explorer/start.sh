#!/bin/bash

hospital_priv_sk_path=$(ls ../crypto-config/peerOrganizations/hospital.com/users/Admin\@hospital.com/msp/keystore/)
patient_priv_sk_path=$(ls ../crypto-config/peerOrganizations/patient.com/users/Admin\@patient.com/msp/keystore/)

# cp -rf ./connection-profile/network_temp.json ./connection-profile/hospital_network.json

sed -i "s/priv_sk/$hospital_priv_sk_path/" ./connection-profile/hospital_network.json
sed -i "s/priv_sk/$patient_priv_sk_path/" ./connection-profile/patient_network.json

docker-compose down -v
docker-compose up -d