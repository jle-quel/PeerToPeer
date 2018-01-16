/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 13:06:16 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func main() {
	// Phase 1
	getHeader := initHeader()
	addPeer := initRoutingTable()
	getHeader().Broadcast()

	fmt.Printf("Id [%s]\n", getHeader().Id)
	fmt.Printf("Addr [%s]\n", getHeader().Addr)
	fmt.Printf("Timestamp [%d]\n\n", getHeader().Timestamp)

	// Phase 2
	go UDPServer(addPeer, getHeader)
	go TCPServer(addPeer)
	go handleSignal(getHeader)
	for {}
}
