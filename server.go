/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 00:23:32 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func debug(peers t_map) {
	for _, value := range peers {
		fmt.Println(value)
	}
	fmt.Printf("\n")
}

/*
**** PUBLIC ********************************************************************
*/

func UDPServer(getHeader func() header, headerCh chan header) {
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		_, err := conn.Read(buf)
		handleErr(err)
		fmt.Println("New peer...")

		peer := decode(buf)
		fmt.Println(peer)
		fmt.Printf("\n")
		go getHeader().Bootstrap(peer.Addr + TCP_PORT)
	}
	conn.Close()
}


func TCPServer(headerCh chan header) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()
	addPeer := initRoutingTable()

	for {
		conn, err := listener.Accept()
		handleErr(err)
		conn.Read(buf)
		fmt.Println("New header...")

		debug(addPeer(decode(buf)))
		conn.Close()
	}
	listener.Close()
}
