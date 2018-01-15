/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 19:18:20 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
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
	go peerServer(addPeer, getHeader)
	go headerServer(addPeer)
	go handleSignal(getHeader)
	for {}
}
