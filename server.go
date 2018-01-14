/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/14 23:58:59 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func UDPServer(getHeader func() header, headerCh chan header) {
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		fmt.Println("Listening for peer...")

		_, err := conn.Read(buf)
		handleErr(err)

		peer := decode(buf)
		go getHeader().Bootstrap(peer.Addr + TCP_PORT)
		// headerCh <- peer
	}
	conn.Close()
}


func TCPServer(headerCh chan header) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()
	addPeer := initRoutingTable()

	for {
		// addPeer(<- headerCh)
		fmt.Println("Listening for header...")

		for { // Time out function
			conn, err := listener.Accept()
			handleErr(err)
			conn.Read(buf)

			temp := decode(buf)
			fmt.Printf("Id [%s]\nAddr [%s]\nTimestamp [%d]\n\n", temp.Id, temp.Addr, temp.Timestamp)
			addPeer(temp)
			conn.Close()
		}
	}
	listener.Close()
}
