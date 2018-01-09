/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listen.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/09 11:46:28 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/09 12:29:42 by jle-quel         ###   ########.fr       */
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

func handlePeer(conn net.Conn) {
	buf := make([]byte, HEADERSIZE)
	conn.Read(buf)
	fmt.Println(string(buf))
}

/*
**** PUBLIC ********************************************************************
*/

func listenTCP() {
	listener, err := net.Listen("tcp", BOOTSTRAP_PORT)

	for {
		fmt.Println("Listening for headers ...")
		handleErr(err)

		conn, err := listener.Accept()
		handleErr(err)
		handlePeer(conn)
		conn.Close()
	}
}
