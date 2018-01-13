/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 18:24:13 by jle-quel         ###   ########.fr       */
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

func initUDPListen() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_PORT)
	handleErr(err)
	conn, err := net.ListenUDP("udp", addr)
	handleErr(err)
	return conn
}

/*
**** PUBLIC ********************************************************************
*/

func UDPServer() {
	fmt.Println("Listening for new peers...")
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		_, err := conn.Read(buf)
		handleErr(err)
		fmt.Println(string(buf))
	}
	conn.Close()
}
