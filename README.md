> 🚀 This project uses Hyperledger Fabric to build the underlying bolckchain network, use go to write smart contract, application layer use gin+fabric-sdk-go, fronted-end use vue+element-ui


## Environment demand

Need the Linux environment with Docker and Docker Compose Docker

Ps Docker install tutorial: [click it](Install.md)


## Deployment

1. clone this project to any content, eg.`/usr/local/fabric-medical`

2. give the project permisson, execute `sudo chmod -R +x /usr/local/fabric-medical/`

3. go into `network` content, execute `./start.sh` start blockchain network

4. go into `chaincode` content, execute `./sc-start.sh` deploy smart contract

5. go into `application` content, execute `./build.sh` compile image, then execute `./start.sh`
   start the application, use explorer to access [http://localhost:8000/web](http://localhost:8000/web)

6. (optional) go into `network/explorer` content, execute `./start.sh` restart blockchain explorer, access [http://localhost:8080](http://localhost:8080), username: admin, password
   123456

## stop or restart

Note: default execute`./start.sh` the all script will be called`./stop.sh`, if you want to stop/restart the project in the case of keeping persistence, please do not execute `./start.sh` repeatedly, The references are as follows:

1. (If start explorer blockchain application) go into `network/explorer` content, execute `docker-compose stop` stop blockchain explorer, execute `docker-compose start`
   start blockchain explorer, execute `docker-compose restart` restart blockchain explorer

2. go into `network` content, execute `docker-compose stop` stop blockchain network, execute `docker-compose start`
   start blockchain network, execute `docker-compose restart` restart blockchain network

3. go into `application` content, the blockchain application is stateless, please execute directly. `./stop.sh`  close blockchain application `./start.sh` restart application

## Clean up environment

Please note that the operation will clean up all data. Please follow this order:

1. (If you started blockchain explorer) go into `network/explorer` content, executable `./stop.sh` close blockchain explorer

2. go into `network` content, execute `./stop.sh` close blockchain network and clean chaincode container

3. go into `application` content, execute `./stop.sh` close blockchain application

## Contents structure

- `application/server` : `fabric-sdk-go` call chaincode(smart contract),`gin` support external access interfaces(RESTful API)


- `application/web` : `vue` + `element-ui` support fronted display pages


- `chaincode` : go (smart contract)


- `network` : Hyperledger Fabric blockchain network configuration

## Function flow


## Presentation effect

<!--
![login]()

![addreal]()

![info]()

![explorer]() 
-->

