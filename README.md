# PeerDiscovery

## Description

This project is a simple implementation of a peer discovery on a p2p network.

The idea behind the project is :
- Every peer is assign a globaly unique ID
- The peer enter a network without knowing anyone on this network and broadcast his arrival
- Every peer will respond to the new peer, with his header
- Then the new peer will launch the network cycle to self repair the network if peer disconnect or else

## Author

Jefferson Le Quellec
