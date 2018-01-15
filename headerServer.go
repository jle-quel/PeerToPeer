/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   headerServer.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 14:12:34 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func headerServer(addPeer func(peer header) t_map) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()

	for {
		fmt.Println("Listening for header...")
		conn, err := listener.Accept()
		switch err {
		case nil:
			conn.Read(buf)
			addPeer(decode(buf))
			conn.Close()
		default:
			return
		}
	}
	listener.Close()
}
