/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 13:55:27 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
)

type t_bytes []byte

/*
**** PRIVATE *******************************************************************
*/

func initServerSocket() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_PORT)
	handleErr(err)
	socket, err := net.ListenUDP("udp", addr)
	handleErr(err)
	return socket
}

func receiveData(socket *net.UDPConn) []byte {
	buf := make([]byte, BUF_SIZE)
	_, _, err := socket.ReadFrom(buf)
	handleErr(err)
	return buf
}

// func cycle(table t_map) {
// 	time.Sleep(42 * time.Second)
// 	broadcast(table.Encode())
// }

/*
**** PUBLIC ********************************************************************
*/

func server(ch chan t_map) {
	socket := initServerSocket()
	addData := routingTable()

	// for {
		buf := t_bytes(receiveData(socket))
		peer := buf.DecodeHeader()
		table := addData(peer)
		broadcast(table.Encode())
		// go cycle(table)
		ch <- table
	// }
	socket.Close()
}
