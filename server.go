/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 20:08:40 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
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

func handlePeer(conn net.Conn) {
	buf := make([]byte, HEADER_SIZE)
	conn.Read(buf)
	fmt.Println(string(buf))
}

/*
**** PUBLIC ********************************************************************
*/

func UDPServer(getHeader func() header) {
	fmt.Println("Listening for new peer...")
	addPeer := initRoutingTable()
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		// Waiting for peer
		_, err := conn.Read(buf)
		handleErr(err)

		// Adding peer to routing table
		peer := decode(buf)
		addPeer(peer)

		// Sending myself to new peer
		getHeader().Send(peer.Addr + TCP_PORT)
	}
	conn.Close()
}


func TCPServer() {
	listener := initTCPListen()

	for {
		fmt.Println("Listening for headers ...")

		conn, err := listener.Accept()
		handleErr(err)
		handlePeer(conn)
		conn.Close()
	}
	listener.Close()
}
