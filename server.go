/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 15:51:47 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"os"
	"time"
	"fmt"
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

func cycle() {
	time.Sleep(1 * time.Second)
	broadcast(s_header{getAddr(), os.Args[1]}.Encode())
}

/*
**** PUBLIC ********************************************************************
*/

func server(ch chan t_map) {
	socket := initServerSocket()
	addData := routingTable()

	for {
		buf := t_bytes(receiveData(socket))
		// broadcast(s_header{getAddr(), os.Args[1]}.Encode())
		peer := buf.DecodeHeader()
		fmt.Println(addData(peer))
	}
	socket.Close()
}
