# PeerDiscovery

## Description

This project is a simple implementation of peer discovery on a p2p network.
It is only working on LAN in ETHERNET/WIFI

The idea behind it is to :
- Learn about network, and how peers can interact with each others
- Experiment with go

When launching, you will be :
- Assign an ID
- Broadcasting the network for PeerDiscovery
- If peer exist on the network, they will send you a tram with `Id, Addr`
- You will be now fully known on the network and you will have your own routingTable
- If you leave the network, you will before leaving telling all peers about your action and every peer will update their routingTable, so the network is self repairing

## Requirements

You need to have installed `Golang`

```
Linux:
sudo apt-get install golang

OSX:
brew install golang

Window:
LOL
```

## Installation and Usage

1. Clone repository
2. Run next command in your terminal :

```
git clone https://github.com/jle-quel/PeerDiscovery.git
cd PeerDiscovery && go build
make && ./21sh
```

## Author

Jefferson Le Quellec
