/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 19:20:30 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
	"bytes"
	"encoding/json"
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

func decode(peer []byte) header {
	var ret header

	b := bytes.NewReader(peer)
	err := json.NewDecoder(b).Decode(&ret)
	handleErr(err)
	return ret
}

func initRoutingTable() func(peer header) t_map {
	m := make(t_map)
	return func(peer header) t_map {
		m[peer.Id] = routingTable{peer.Addr}
		return m
	}
}

/*
**** PUBLIC ********************************************************************
*/

func UDPServer() {
	fmt.Println("Listening for new peers...")
	addPeer := initRoutingTable()
	buf := make([]byte, HEADER_SIZE)
	conn := initUDPListen()

	for {
		_, err := conn.Read(buf)
		handleErr(err)
		fmt.Println(addPeer(decode(buf)))
	}
	conn.Close()
}
