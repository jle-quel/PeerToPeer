/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 19:17:43 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"encoding/json"
	"bytes"
	"fmt"
)

/*
**** PRIVATE *******************************************************************
*/

func decodeData(buf []byte) s_header {
	var peer s_header

	b := bytes.NewReader(buf)
	json.NewDecoder(b).Decode(&peer)
	return peer
}

/*
**** PUBLIC ********************************************************************
*/

func server() {
	var table routingTable

	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_PORT)
	handleErr(err)
	socket, err := net.ListenUDP("udp", addr)
	handleErr(err)

	for {
		buf := make([]byte, BUF_SIZE)
		_, _, err := socket.ReadFrom(buf)
		handleErr(err)

		peer := decodeData(buf)
		f := addData(peer)
	}
	socket.Close()
}
