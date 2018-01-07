/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listen.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 14:12:10 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/07 21:56:58 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	// "net"
	"fmt"
)

/*
**** PRIVATE *******************************************************************
*/

// func initListenConn() *net.UDPConn{
// 	addr, err := net.ResolveUDPAddr("udp", PORT_ADDR)
// 	handleErr(err)
// 	conn, err := net.ListenUDP("udp", addr)
// 	handleErr(err)
// 	return conn
// }
//
// func receivePeer() string {
// 	buf := make([]byte, 64)
// 	conn := initListenConn()
// 	_, _, err := conn.ReadFrom(buf)
// 	handleErr(err)
// 	return string(buf)
// }

/*
**** PUBLIC ********************************************************************
*/

func listen() {
	// var table t_map
	// addData := tableClosure()

	for {
		//Wait for peers
		fmt.Println("Listening for new peers ...")
		// buf := receivePeer()

		//Add data to t_map
		// table = addData(getAddr(addr) , buf)

		//If nearest peer send him t_map

		// handleErr(err)
		// conn.Close()
	}
}
