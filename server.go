/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 16:44:17 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 16:45:16 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
	"encoding/json"
	"bytes"
)

/*
**** PUBLIC ********************************************************************
*/

func decodeData(buf []byte) s_header {
	var peer s_header

	b := bytes.NewReader(buf)
	json.NewDecoder(b).Decode(&peer)
	return peer
}

func server() {

	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_PORT)
	handleErr(err)
	conn, err := net.ListenUDP("udp", addr)
	handleErr(err)
	buf := make([]byte, BUF_SIZE)
	_, _, err = conn.ReadFrom(buf)
	handleErr(err)

	fmt.Println(decodeData(buf))
	conn.Close()
}
