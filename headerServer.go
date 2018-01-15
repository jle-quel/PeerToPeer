/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   headerServer.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 14:16:31 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func debug(peers t_map) {
	fmt.Println("Routing Table")
	for _, value := range peers {
		fmt.Println(value)
	}
	fmt.Printf("\n")
}

/*
**** PUBLIC ********************************************************************
*/

func headerServer(addPeer func(peer header) t_map) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()

	for {
		fmt.Println("Listening for header...")
		conn, _ := listener.Accept()
		conn.Read(buf)
		debug(addPeer(decode(buf)))
		conn.Close()
	}
	listener.Close()
}
