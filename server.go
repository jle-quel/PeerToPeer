/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 20:57:43 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"encoding/json"
	"bytes"
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

func receiveData(socket *net.UDPConn) []byte {
	buf := make([]byte, BUF_SIZE)
	_, _, err := socket.ReadFrom(buf)
	handleErr(err)
	return buf
}

func decodeData(buf []byte) s_header {
	var peer s_header

	b := bytes.NewReader(buf)
	json.NewDecoder(b).Decode(&peer)
	return peer
}

/*
**** PUBLIC ********************************************************************
*/

func server(ch chan t_map) {
	socket := initServerSocket()
	addData := routingTable()

	for {
		buf := receiveData(socket)
		peer := decodeData(buf)
		ch <- addData(peer)
	}
	socket.Close()
}
