/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   peerServer.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/15 10:58:36 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 16:49:30 by jle-quel         ###   ########.fr       */
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

func handleConn(conn *net.UDPConn) header {
	buf := make([]byte, HEADER_SIZE)

	_, err := conn.Read(buf)
	handleErr(err)
	peer := decode(buf)
	return peer
}

/*
**** PUBLIC ********************************************************************
*/

func peerServer(addPeer func(peer header) (t_map, int), getHeader func() header) {
	conn := initUDPListen()

	for {
		fmt.Println("Listening for peer...")
		peer := handleConn(conn)
		routingTable, _ := addPeer(peer)
		// switch err {
		// case 0:
		// 	getHeader().Bootstrap(peer.Addr + TCP_PORT)
		// }
		debug(routingTable, 42)
	}
	conn.Close()
}
