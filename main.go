/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 13:02:07 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"bufio"
	"os"
)

/*
**** PUBLIC ********************************************************************
*/

func loop(ch chan t_map) {
	for {
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadByte()
		handleErr(err)

		switch ret {
		case 49:
			fmt.Println(<- ch)
		default:
			fmt.Println("not yet")
		}
	}
}

func main() {
	// Phase 1
	ch := make(chan t_map)
	getHeader := initHeader()
	addPeer := initRoutingTable()
	getHeader().Broadcast()

	fmt.Printf("Id [%s]\n", getHeader().Id)
	fmt.Printf("Addr [%s]\n", getHeader().Addr)
	fmt.Printf("Timestamp [%d]\n\n", getHeader().Timestamp)

	// Phase 2
	go UDPServer(addPeer, getHeader, ch)
	go TCPServer(addPeer, ch)
	go handleSignal(getHeader)
	loop(ch)
}
