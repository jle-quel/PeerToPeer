/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   connection.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/09 12:06:08 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/09 15:11:01 by jle-quel         ###   ########.fr       */
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

func initUDPConn() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_ADDR + BROADCAST_PORT)
	handleErr(err)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return conn
}

func initTCPConn(ip string) *net.TCPConn {
	addr, err := net.ResolveTCPAddr("tcp", ip + BOOTSTRAP_PORT)
	handleErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
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
	handleErr(err)
	conn.Close()
}

func peerBootstrap(peer []byte, ip string) {
	fmt.Println("Bootstraping ...")

	conn := initTCPConn(ip)
	_, err := conn.Write(peer)
	handleErr(err)
	conn.Close()
}
