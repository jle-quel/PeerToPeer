/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 19:34:35 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 13:16:10 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
)

/*
**** PUBLIC ********************************************************************
*/

func initUDPListen() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_PORT)
	handleErr(err)
	conn, err := net.ListenUDP("udp", addr)
	handleErr(err)
	return conn
}

func initUDPConn() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_ADDR)
	handleErr(err)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return conn
}

func initTCPConn(ip string) *net.TCPConn {
	addr, err := net.ResolveTCPAddr("tcp", ip)
	handleErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	handleErr(err)
	return conn
}

func initTCPListen() *net.TCPListener {
	addr, err := net.ResolveTCPAddr("tcp", TCP_PORT)
	handleErr(err)
	listen, err := net.ListenTCP("tcp", addr)
	handleErr(err)
	return listen
}

func initRoutingTable() func(peer header) (t_map, int) {
	var ret int
	m := make(t_map)
	return func(peer header) (t_map, int) {
		_, ok := m[peer.Id]
		switch ok {
		case true:
			delete(m, peer.Id)
			ret = 1
		case false:
			m[peer.Id] = routingTable{peer.Addr}
			ret = 0
		}
		return m, ret
	}
}
