/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/14 19:17:30 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/*
**** PUBLIC ********************************************************************
*/

func main() {
	// Peer Discovery
	getHeader := initHeader()
	// getHeader().SendUDP()
	// broadcast(getHeader().Encode())
	// go UDPServer(getHeader)
	TCPServer()
	getHeader().Bootstrap("192.168.0.11" + TCP_PORT)
}
