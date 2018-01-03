/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/02 13:30:49 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 20:58:39 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"os"
	"encoding/json"
	"bytes"
)

type s_header struct {
	Addr string
	Username string
	Hash uint32
}

/*
**** PRIVATE *******************************************************************
*/

func initBroadcastSocket() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_ADDR)
	handleErr(err)
	socket, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return socket
}

func encodeData(peer s_header) *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(peer)
	return buf
}

/*
**** PUBLIC ********************************************************************
*/

func broadcast() {
	socket := initBroadcastSocket()
	buf := encodeData(s_header{getAddr(), os.Args[1], getHash(os.Args[1])})
	socket.Write(buf.Bytes())
	socket.Close()
}
