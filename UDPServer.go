/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   UDPServer.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/15 10:58:36 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 13:01:46 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"fmt"
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

func UDPServer(addPeer func(peer header) (t_map, int), getHeader func() header, ch chan t_map) {
	conn := initUDPListen()

	for {
		fmt.Println("Listening for UDPConn...")
		peer := handleConn(conn)
		routingTable, err := addPeer(peer)
		switch err {
		case 0:
			go getHeader().Bootstrap(peer.Addr + TCP_PORT)
		}
		ch <- routingTable
	}
	conn.Close()
}
