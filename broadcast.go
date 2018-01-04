/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/02 13:30:49 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 14:20:20 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"bytes"
	"fmt"
)

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

/*
**** PUBLIC ********************************************************************
*/

func broadcast(data *bytes.Buffer) {
	socket := initBroadcastSocket()
	size, err := socket.Write(data.Bytes())
	handleErr(err)
	fmt.Println(size, string(data.Bytes()))
	socket.Close()
}
