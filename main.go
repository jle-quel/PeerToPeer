/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/02 19:51:17 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
	"os"
)

/*
**** PRIVATE *******************************************************************
*/

func server() {

	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_PORT)
	handleErr(err)
	conn, err := net.ListenUDP("udp", addr)
	handleErr(err)
	defer conn.Close()

	buf := make([]byte, 1)
	size, _, err := conn.ReadFrom(buf)
	handleErr(err)
	fmt.Println(size)
}

/*
**** PUBLIC ********************************************************************
*/

func main() {

	switch len(os.Args) {
	case 2:
		fmt.Println("Broadcasting ...")
		broadcast()
	default:
		fmt.Println("Waiting for peers ...")
		server()
	}
}
