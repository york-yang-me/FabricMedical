#!/bin/bash

hospital_priv_sk_path=$(ls ../crypto-config/peerOrganizations/hospital.com/users/Admin\@hospital.com/msp/keystore/)
institute_priv_sk_path=$(ls ../crypto-config/peerOrganizations/institute.com/users/Admin\@institute.com/msp/keystore/)

# cp -rf ./connection-profile/network_temp.json ./connection-profile/hospital_network.json

sed -i "s/priv_sk/$hospital_priv_sk_path/" ./connection-profile/hospital_network.json
sed -i "s/priv_sk/$institute_priv_sk_path/" ./connection-profile/institute_network.json

docker-compose down -v
docker-compose up -d