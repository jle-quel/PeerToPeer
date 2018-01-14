/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/14 19:09:31 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PRIVATE *******************************************************************
*/

func initRoutingTable() func(peer header) t_map {
	m := make(t_map)
	return func(peer header) t_map {
		m[peer.Id] = routingTable{peer.Addr}
		return m
	}
}

/*
**** PUBLIC ********************************************************************
*/

func UDPServer(getHeader func() header) {
	addPeer := initRoutingTable()
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		fmt.Println("Listening for new peer...")

		// Waiting for peer
		_, err := conn.Read(buf)
		handleErr(err)

		// Adding peer to routing table
		peer := decode(buf)
		addPeer(peer)

		// Sending myself to new peer
		getHeader().Bootstrap(peer.Addr + TCP_PORT)
	}
	conn.Close()
}


func TCPServer() {
	listener := initTCPListen()
	buf := make([]byte, HEADER_SIZE)

	for {
		fmt.Println("Listening for headers...")

		conn, err := listener.Accept()
		handleErr(err)
		conn.Read(buf)

		peer := decode(buf)
		fmt.Printf("Id [%s]\nAddr [%s]\nTimestamp [%d]\n\n", peer.Id, peer.Addr, peer.Timestamp)
		
		conn.Close()
	}
	listener.Close()
}
