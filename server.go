/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 19:26:12 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

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

func receiveData(socket *net.UDPConn) t_bytes {
	header := make([]byte, HEADER_SIZE)
	_, _, err := socket.ReadFrom(header)
	handleErr(err)
	return t_bytes(header)
}

/*
**** PUBLIC ********************************************************************
*/

func server() {
	tableClosure := routingTable()

	for {
		socket := initServerSocket()
		fmt.Println("Waiting for peers ...")
		header := receiveData(socket)
		fmt.Println(tableClosure(header.DecodeHeader()))
		socket.Close()
		fmt.Println("Broadcasting ...")
		broadcast(s_header{getAddr(), os.Args[1]}.Encode())
	}
}
