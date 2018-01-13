/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 10:41:07 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 11:48:41 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"bytes"
	"net"
	"encoding/json"
)

/*
**** PRIVATE *******************************************************************
*/

func initUDPConn() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_ADDR)
	handleErr(err)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return conn
}

/*
**** PUBLIC ********************************************************************
*/

func (self header) Encode() []byte {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(self)
	handleErr(err)
	return b.Bytes()
}

func broadcast(b []byte) {
	fmt.Println("Broadcasting", string(b))

	conn := initUDPConn()
	_, err := conn.Write(b)
	handleErr(err)
	conn.Close()
}
