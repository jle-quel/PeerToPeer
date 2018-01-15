# PeerDiscovery

## Description

This project is a simple implementation of peer discovery on a p2p network.
It is only working on LAN in ETHERNET/WIFI

The idea behind it is to :
- Learn about network, and how peers can interact with each others in a fully decentralized architecture
- Experiment with go

## Actions

When launching, you will be :
- Assign an ID
- Broadcasting the network for PeerDiscovery
- If peer exist on the network, they will send you a tram with `Id, Addr`
- You will be now fully known on the network and you will have your own routingTable
- If you leave the network. You will, before leaving, telling all peers about your action and every peer will update their routingTable, so the network is self repairing

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
./PeerDiscovery
```

## Contributing

    Fork this Repo first
    Clone your Repo
    Install dependencies by $ npm install
    Checkout a feature branch
    Feel free to add your features
    Make sure your features are fully tested
    Publish your local branch, Open a pull request
    Enjoy hacking <3

## MIT license

### Copyright (c) 2017 jle-quel

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
