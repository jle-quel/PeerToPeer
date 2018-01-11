/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listen.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/09 11:46:28 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/10 14:29:32 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"fmt"
	_ "time"
)

/*
**** PRIVATE *******************************************************************
*/

// func handlePeer(conn net.Conn, ch chan s_header) {
// 		buf := make([]byte, HEADERSIZE)
// 		conn.Read(buf)
// 		peer := t_byte(buf).Decode()
//
// 		if peer.Maj > 1515510300 {
// 			ch <- peer
// 		} else {
// 			close(ch)
// 		}
// }

func initListenerTCP() *net.TCPListener {
	addr, err := net.ResolveTCPAddr("tcp", BOOTSTRAP_PORT)
	handleErr(err)
	listener, err := net.ListenTCP("tcp", addr)
	handleErr(err)
	return listener
}

/*
**** PUBLIC ********************************************************************
*/

func listenTCP() {
	listener := initListenerTCP()

	for {
		fmt.Println("Listening for headers ...")

		conn, err := listener.Accept()
		handleErr(err)

		// go handlePeer(conn, ch)
		// fmt.Println(<- ch)
		conn.Close()
	}
}
