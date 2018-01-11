# PeerDiscovery

## Description

This project is a simple implementation of peer discovery on a p2p network.

The idea behind the project is :
- The peer enter a network without knowing anyone on this network and broadcast his arrival
- Every peer will respond to the new peer, with his header
- Then the new peer will launch the network cycle to self repair the network if peer disconnect or else

The network is scalable for a mid-network and works only, for now, on local network ETHERNET/WIFI

## Author

Jefferson Le Quellec
