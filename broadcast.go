/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 11:36:52 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/07 21:54:39 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"fmt"
	"bytes"
)

/*
**** PRIVATE *******************************************************************
*/

func initBroadcastConn() *net.UDPConn{
	addr, err := net.ResolveUDPAddr("udp", BROADCAST_ADDR)
	handleErr(err)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)
	return conn
}

/*
**** PUBLIC ********************************************************************
*/

func broadcast(data *bytes.Buffer) {
	conn := initBroadcastConn()
	fmt.Println("Broadcasting ...")
	_, err := conn.Write(data.Bytes())
	fmt.Println(string(data.Bytes()))
	handleErr(err)
	conn.Close()
}
