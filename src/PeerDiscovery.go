/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   peerDiscovery.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 22:00:34 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/08 22:28:51 by jle-quel         ###   ########.fr       */
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

func initUDPConn() *net.UDPConn{
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_ADDR + BROADCAST_PORT)
	handleErr(err)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return conn
}

/*
**** PUBLIC ********************************************************************
*/

func peerDiscovery(peer []byte) {
	fmt.Println("Broadcasting ...")

	conn := initUDPConn()
	_, err := conn.Write(peer)
	fmt.Println(string(peer))
	handleErr(err)
	conn.Close()
}
