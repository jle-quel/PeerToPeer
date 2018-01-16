/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   TCPServer.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 12:58:21 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func TCPServer(addPeer func(peer header) (t_map, int), ch chan t_map) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()

	for {
		fmt.Println("Listening for TCPConn...")
		conn, _ := listener.Accept()
		conn.Read(buf)
		routingTable, _ := addPeer(decode(buf))
		ch <- routingTable
		conn.Close()
	}
	listener.Close()
}
